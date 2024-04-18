package tests

import (
	pairHandlers "GophKeeper/app/internal/adapters/http/pair/handlers"
	"GophKeeper/app/internal/entity/pairs"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPairService := mock_pairhandlers.mock_pairhandlers.NewMockpairService(ctrl)
	mockLogger := logger{}

	pairHandlers := pairHandlers.NewPairHandler(mockLogger, mockPairService)
	c := context.Background()

	//===========================TEST 1===========================
	mockPairService.EXPECT().Delete(c, "pair1", 1).Return(nil)

	//===========================TEST 2===========================
	mockPairService.EXPECT().Delete(c, "pair1", 1).Return(pairs.ErrPairNothingToDelete)

	//===========================TEST 3===========================
	mockPairService.EXPECT().Delete(c, "pair1", 1).Return(errors.New("new error"))

	//===========================TEST 4===========================
	// in this test there is no call to the pairService

	testCase := []struct {
		name           string
		statusCode     int
		expectedBody   string
		valueInContext string
		urlKey         string
	}{
		{
			name:           "test_1",
			statusCode:     http.StatusOK,
			expectedBody:   `"ok"`,
			valueInContext: "userID",
			urlKey:         "pair1",
		},
		{
			name:           "test_3",
			statusCode:     http.StatusNoContent,
			expectedBody:   `"no content"`,
			valueInContext: "userID",
			urlKey:         "pair1",
		},
		{
			name:           "test_3",
			statusCode:     http.StatusInternalServerError,
			expectedBody:   `"internal server error"`,
			valueInContext: "userID",
			urlKey:         "pair1",
		},
		{
			name:           "test_4",
			statusCode:     http.StatusInternalServerError,
			expectedBody:   `"internal server error"`,
			valueInContext: "ID_OF_USER",
			urlKey:         "pair1",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(
				http.MethodPost,
				"/someURL",
				nil,
			)

			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			ctx.Set(tc.valueInContext, 1)
			ctx.SetParamNames("name")
			ctx.SetParamValues(tc.urlKey)

			err := pairHandlers.Delete(ctx)
			assert.NoError(t, err, "Error creating test request")

			assert.Equal(t, tc.statusCode, rec.Code)
			assert.Equal(t, tc.expectedBody, strings.TrimSuffix(rec.Body.String(), "\n"))
		})
	}
}
