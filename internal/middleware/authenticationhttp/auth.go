package authenticationhttp

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"GophKeeper/internal/jwt"

	"github.com/labstack/echo/v4"
)

const (
	internalServerError = "internal server error"
	invalidToken        = "your token is invalid"
	expireToken         = "token has expired"
)

type logger interface {
	Info(msg string)
	Error(e error)
}

type tokenManager interface {
	GetUserID(ctx context.Context, tokenString string) (int, error)
}

type AuthMiddleware struct {
	tokenManager tokenManager
	logger       logger
}

func NewAuthMiddleware(t tokenManager, l logger) *AuthMiddleware {
	return &AuthMiddleware{
		logger:       l,
		tokenManager: t,
	}
}

func (a *AuthMiddleware) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, jwt.ErrInvalidToken)
		}
		token, err := splitToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, invalidToken)
		}

		userID, err := a.tokenManager.GetUserID(c.Request().Context(), token)
		if err != nil {
			a.logger.Error(err)
			if errors.Is(err, jwt.ErrInvalidToken) {
				return echo.NewHTTPError(http.StatusUnauthorized, invalidToken)
			}

			if errors.Is(err, jwt.ErrExpiredToken) {
				return echo.NewHTTPError(http.StatusUnauthorized, expireToken)
			}
			return echo.NewHTTPError(http.StatusInternalServerError, internalServerError)
		}

		c.Set("UserID", userID)
		return next(c)
	}
}

func splitToken(token string) (string, error) {
	arr := strings.Split(token, " ")
	if len(arr) != 3 {
		return "", jwt.ErrInvalidToken
	}

	if arr[1] != "Bearer" {
		return "", jwt.ErrInvalidToken
	}

	if strings.Count(arr[2], ".") != 2 {
		return "", jwt.ErrInvalidToken
	}

	return arr[2], nil
}
