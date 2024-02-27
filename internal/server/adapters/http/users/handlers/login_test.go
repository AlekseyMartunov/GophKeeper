package userhandlers

import (
	mock_userhandlers "GophKeeper/internal/server/adapters/http/users/handlers/mock"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"GophKeeper/internal/server/entity/users"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type myMockLogger struct{}

func (l myMockLogger) Error(e error) {}
func (l myMockLogger) Info(s string) {}

func TestLoginHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := myMockLogger{}
	mockUserService := mock_userhandlers.NewMockuserService(ctrl)
	mockJWT := mock_userhandlers.NewMocktokenJWTManager(ctrl)

	userHandlers := NewUserHandler(mockUserService, mockLogger, mockJWT)

	//===========================TEST 1===========================
	u1 := users.User{
		Password: "pass",
		Login:    "123",
	}
	mockUserService.EXPECT().GetExternalID(context.Background(), u1).Return("ID", nil)
	mockJWT.EXPECT().CreateToken("ID").Return("xxxxxxxxxxxxx", nil)

	//===========================TEST 2===========================
	// in this test there is no call to the userService
	// since the handler completes the function before

	//===========================TEST 3===========================
	u3 := users.User{
		Password: "pass",
		Login:    "123",
	}
	mockUserService.EXPECT().GetExternalID(context.Background(), u3).Return(
		"",
		users.ErrUserDoseNotExist,
	)

	//===========================TEST 4===========================
	u4 := users.User{
		Password: "pass",
		Login:    "123",
	}
	mockUserService.EXPECT().GetExternalID(context.Background(), u4).Return(
		"",
		errors.New("internal new error"),
	)

	//===========================TEST 5===========================
	u5 := users.User{
		Password: "pass",
		Login:    "123",
	}
	mockUserService.EXPECT().GetExternalID(context.Background(), u5).Return("ID", nil)
	mockJWT.EXPECT().CreateToken("ID").Return("", errors.New("some error"))

	testCase := []struct {
		name         string
		body         string
		statusCode   int
		bodyResponse string
	}{
		{
			name:         "test_1",
			body:         `{"login":"123", "password":"pass"}`,
			statusCode:   200,
			bodyResponse: `"Your Token: Bearer xxxxxxxxxxxxx"`,
		},
		{
			name:         "test_2",
			body:         `{"login 123", "password":"pass"`,
			statusCode:   400,
			bodyResponse: `"the request form is incorrect or the request does not contain the required field"`,
		},
		{
			name:         "test_3",
			body:         `{"login":"123", "password":"pass"}`,
			statusCode:   204,
			bodyResponse: `"user with this login and password dose not exists"`,
		},
		{
			name:         "test_4",
			body:         `{"login":"123", "password":"pass"}`,
			statusCode:   500,
			bodyResponse: `"internal server error"`,
		},
		{
			name:         "test_5",
			body:         `{"login":"123", "password":"pass"}`,
			statusCode:   500,
			bodyResponse: `"internal server error"`,
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

			err := userHandlers.Login(ctx)

			assert.NoError(t, err, "Error creating test request")

			assert.Equal(t, tc.statusCode, rec.Code)
			assert.Equal(t, tc.bodyResponse, strings.TrimSuffix(rec.Body.String(), "\n"))
		})
	}
}
