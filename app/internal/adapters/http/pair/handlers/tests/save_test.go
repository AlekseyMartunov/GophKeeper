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

type logger struct{}

func (l logger) Info(s string) {}
func (l logger) Error(r error) {}

func TestSave(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPairService := mock_pairhandlers.mock_pairhandlers.NewMockpairService(ctrl)
	mockLogger := logger{}

	pairHandlers := pairHandlers.NewPairHandler(mockLogger, mockPairService)
	c := context.Background()

	//===========================TEST 1===========================
	p1 := pairs.Pair{
		Login:    "l",
		Password: "p",
		Name:     "p1",
		UserID:   1,
	}
	mockPairService.EXPECT().Save(c, p1).Return(nil)

	//===========================TEST 2===========================
	// in test 2 handler return error before pairService called

	//===========================TEST 3===========================
	p3 := pairs.Pair{
		Login:    "l",
		Password: "p",
		Name:     "p1",
		UserID:   1,
	}
	mockPairService.EXPECT().Save(c, p3).Return(pairs.ErrPairAlreadyExists)

	//===========================TEST 4===========================
	// in test 4 handler return error before pairService called

	//===========================TEST 5===========================
	p5 := pairs.Pair{
		Login:    "l",
		Password: "p",
		Name:     "p1",
		UserID:   1,
	}
	mockPairService.EXPECT().Save(c, p5).Return(errors.New("new err"))

	testCase := []struct {
		name           string
		body           string
		statusCode     int
		expectedBody   string
		valueInContext string
	}{
		{
			name:           "test_1",
			body:           `{"login": "l", "password": "p", "name": "p1"}`,
			statusCode:     http.StatusOK,
			expectedBody:   `"ok"`,
			valueInContext: "userID",
		},
		{
			name:           "test_2",
			body:           `"login": "l", password": "p", name": "p1"}`,
			statusCode:     http.StatusBadRequest,
			expectedBody:   `"the request form is incorrect or the request does not contain the required field"`,
			valueInContext: "userID",
		},
		{
			name:           "test_3",
			body:           `{"login": "l", "password": "p", "name": "p1"}`,
			statusCode:     http.StatusConflict,
			expectedBody:   `"you already have pair with this name"`,
			valueInContext: "userID",
		},
		{
			name:           "test_4",
			body:           `{"login": "l", "password": "p", "name": "p1"}`,
			statusCode:     http.StatusInternalServerError,
			expectedBody:   `"internal server error"`,
			valueInContext: "userName",
		},
		{
			name:           "test_5",
			body:           `{"login": "l", "password": "p", "name": "p1"}`,
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

			err := pairHandlers.Save(ctx)
			assert.NoError(t, err, "Error creating test request")

			assert.Equal(t, tc.statusCode, rec.Code)
			assert.Equal(t, tc.expectedBody, strings.TrimSuffix(rec.Body.String(), "\n"))
		})
	}
}
