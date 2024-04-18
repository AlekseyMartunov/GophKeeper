package userhandlers

import (
	"GophKeeper/app/internal/entity/users"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := myMockLogger{}
	mockUserService := mock_userhandlers.mock_userhandlers.NewMockuserService(ctrl)
	mockJWT := mock_userhandlers.mock_userhandlers.NewMocktokenJWTManager(ctrl)

	userHandlers := NewUserHandler(mockUserService, mockLogger, mockJWT)

	//===========================TEST 1===========================

	u1 := users.User{
		Password: "pass",
		Login:    "123",
	}
	mockUserService.EXPECT().Save(context.Background(), u1).Return(nil)

	//===========================TEST 2===========================
	// in this test there is no call to the userService
	// since the handler completes the function before

	//===========================TEST 3===========================
	u3 := users.User{
		Password: "pass",
		Login:    "123",
	}
	mockUserService.EXPECT().Save(context.Background(), u3).Return(users.ErrUserAlreadyExists)

	//===========================TEST 4===========================
	u4 := users.User{
		Password: "pass",
		Login:    "123",
	}
	err := errors.New("some internal error")
	mockUserService.EXPECT().Save(context.Background(), u4).Return(err)

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
			bodyResponse: `{"login":"123","password":"pass"}`,
		},
		{
			name:         "test_2",
			body:         `{"login":"123", "password "pass"`,
			statusCode:   400,
			bodyResponse: `"the request form is incorrect or the request does not contain the required field"`,
		},
		{
			name:         "test_3",
			body:         `{"login":"123", "password":"pass"}`,
			statusCode:   409,
			bodyResponse: `"this login already used by another user"`,
		},
		{
			name:         "test_4",
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

			err := userHandlers.Register(ctx)

			assert.NoError(t, err, "Error creating test request")

			assert.Equal(t, tc.statusCode, rec.Code)
			assert.Equal(t, tc.bodyResponse, strings.TrimSuffix(rec.Body.String(), "\n"))
		})
	}
}
