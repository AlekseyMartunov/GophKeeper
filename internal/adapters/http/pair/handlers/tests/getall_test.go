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
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPairService := mock_pairhandlers.NewMockpairService(ctrl)
	mockLogger := logger{}

	pairHandlers := pairHandlers.NewPairHandler(mockLogger, mockPairService)
	c := context.Background()

	timeForTest, err := time.Parse("Mon Jan 02 15:04:05 -0700 2006", "Tue Sep 09 12:45:00 +0000 2022")
	if err != nil {
		assert.NoError(t, err)
	}

	//===========================TEST 1===========================
	p1 := []pairs.Pair{
		{
			Name:        "pair for example.com1",
			CreatedTime: timeForTest,
		},
		{
			Name:        "pair for example.com2",
			CreatedTime: timeForTest,
		},
	}
	mockPairService.EXPECT().GetAll(c, 1).Return(p1, nil)

	//===========================TEST 2===========================
	// in test 2 handler return error before pairService called

	//===========================TEST 3===========================
	mockPairService.EXPECT().GetAll(c, 1).Return(nil, pairs.ErrPairDoseNotExist)

	//===========================TEST 4===========================
	mockPairService.EXPECT().GetAll(c, 1).Return(nil, errors.New("new err"))

	testCase := []struct {
		name           string
		body           string
		statusCode     int
		expectedBody   string
		valueInContext string
	}{
		{
			name:           "test_1",
			body:           ``,
			statusCode:     http.StatusOK,
			expectedBody:   `[{"name":"pair for example.com1","created_time":"2022-09-09T12:45:00Z"},{"name":"pair for example.com2","created_time":"2022-09-09T12:45:00Z"}]`,
			valueInContext: "userID",
		},
		{
			name:           "test_2",
			body:           ``,
			statusCode:     http.StatusInternalServerError,
			expectedBody:   `"internal server error"`,
			valueInContext: "userNAME",
		},
		{
			name:           "test_3",
			body:           ``,
			statusCode:     http.StatusNoContent,
			expectedBody:   `"you do not have any pairs"`,
			valueInContext: "userID",
		},
		{
			name:           "test_4",
			body:           ``,
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

			err := pairHandlers.GetAll(ctx)
			assert.NoError(t, err, "Error creating test request")

			assert.Equal(t, tc.statusCode, rec.Code)
			assert.Equal(t, tc.expectedBody, strings.TrimSuffix(rec.Body.String(), "\n"))
		})
	}
}
