package tests

import (
	pairHandlers "GophKeeper/internal/adapters/http/pair/handlers"
	"GophKeeper/internal/adapters/http/pair/handlers/mocks"
	"GophKeeper/internal/entity/pairs"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPairService := mock_pairhandlers.NewMockpairService(ctrl)
	mockLogger := logger{}

	pairHandlers := pairHandlers.NewPairHandler(mockLogger, mockPairService)
	c := context.Background()

	//===========================TEST 1===========================
	mockPairService.EXPECT().Delete(c, "pair1", 1).Return(nil)

	//===========================TEST 2===========================
	// in test 2 handler return error before pairService called

	//===========================TEST 3===========================
	// in test 3 handler return error before pairService called

	//===========================TEST 4===========================
	mockPairService.EXPECT().Delete(c, "pair1", 1).Return(pairs.ErrPairNothingToDelete)

	//===========================TEST 5===========================
	mockPairService.EXPECT().Delete(c, "pair1", 1).Return(errors.New("new error"))

	testCase := []struct {
		name           string
		body           string
		statusCode     int
		expectedBody   string
		valueInContext string
	}{
		{
			name:           "test_1",
			body:           `{"name":"pair1"}`,
			statusCode:     http.StatusOK,
			expectedBody:   `"ok"`,
			valueInContext: "userID",
		},
		{
			name:           "test_2",
			body:           `"name":"pair1"}`,
			statusCode:     http.StatusBadRequest,
			expectedBody:   `"the request form is incorrect or the request does not contain the required field"`,
			valueInContext: "userID",
		},
		{
			name:           "test_3",
			body:           `{"name":"pair1"}`,
			statusCode:     http.StatusInternalServerError,
			expectedBody:   `"internal server error"`,
			valueInContext: "userNAME",
		},
		{
			name:           "test_4",
			body:           `{"name":"pair1"}`,
			statusCode:     http.StatusNoContent,
			expectedBody:   `"no content"`,
			valueInContext: "userID",
		},
		{
			name:           "test_5",
			body:           `{"name":"pair1"}`,
			statusCode:     http.StatusInternalServerError,
			expectedBody:   `"internal server error"`,
			valueInContext: "userID",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(
				http.MethodPost,
				"/someURL",
				strings.NewReader(tc.body))

			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			ctx.Set(tc.valueInContext, 1)

			err := pairHandlers.Delete(ctx)
			assert.NoError(t, err, "Error creating test request")

			assert.Equal(t, tc.statusCode, rec.Code)
			assert.Equal(t, tc.expectedBody, strings.TrimSuffix(rec.Body.String(), "\n"))
		})
	}
}
