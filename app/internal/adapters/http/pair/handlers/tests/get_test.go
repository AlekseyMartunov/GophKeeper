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
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPairService := mock_pairhandlers.mock_pairhandlers.NewMockpairService(ctrl)
	mockLogger := logger{}

	pairHandlers := pairHandlers.NewPairHandler(mockLogger, mockPairService)
	c := context.Background()

	timeForTest, err := time.Parse("Mon Jan 02 15:04:05 -0700 2006", "Tue Sep 09 12:45:00 +0000 2022")
	if err != nil {
		assert.NoError(t, err)
	}

	//===========================TEST 1===========================
	p1 := pairs.Pair{
		Login:       "some log",
		Password:    "some pass",
		Name:        "pair for example.com",
		CreatedTime: timeForTest,
	}
	mockPairService.EXPECT().Get(c, "userName_1", 1).Return(p1, nil)

	//===========================TEST 2===========================
	mockPairService.EXPECT().Get(c, "userName_1", 1).Return(pairs.Pair{}, pairs.ErrPairDoseNotExist)

	//===========================TEST 3===========================
	mockPairService.EXPECT().Get(c, "userName_1", 1).Return(pairs.Pair{}, errors.New("new err"))

	//===========================TEST 4===========================
	// in test 4 handler return error before pairService called

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
			expectedBody:   `{"login":"some log","password":"some pass","name":"pair for example.com","created_time":"2022-09-09T12:45:00Z"}`,
			valueInContext: "userID",
			urlKey:         "userName_1",
		},
		{
			name:           "test_2",
			statusCode:     http.StatusNoContent,
			expectedBody:   `"you do not have pair with this name"`,
			valueInContext: "userID",
			urlKey:         "userName_1",
		},
		{
			name:           "test_3",
			statusCode:     http.StatusInternalServerError,
			expectedBody:   `"internal server error"`,
			valueInContext: "userID",
			urlKey:         "userName_1",
		},
		{
			name:           "test_4",
			statusCode:     http.StatusInternalServerError,
			expectedBody:   `"internal server error"`,
			valueInContext: "ID_OF_USER",
			urlKey:         "userName_1",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(
				http.MethodGet,
				"/someURL",
				nil,
			)

			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetParamNames("name")
			ctx.SetParamValues(tc.urlKey)

			ctx.Set(tc.valueInContext, 1)

			err = pairHandlers.Get(ctx)
			assert.NoError(t, err, "Error creating test request")

			assert.Equal(t, tc.statusCode, rec.Code)
			assert.Equal(t, tc.expectedBody, strings.TrimSuffix(rec.Body.String(), "\n"))
		})
	}
}
