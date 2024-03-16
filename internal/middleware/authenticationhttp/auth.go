package authenticationhttp

import (
	"context"
	"errors"
	"net/http"
	"strings"

	tokenPackage "GophKeeper/internal/entity/token"

	"github.com/labstack/echo/v4"
)

const (
	internalServerError = "internal server error"
	invalidToken        = "your token is invalid"
	expireToken         = "token has expired"
	tokenIsBlocked      = "token is blocked"
	noTokenFound        = "no such token exists"
)

type logger interface {
	Info(msg string)
	Error(e error)
}

type tokenService interface {
	GetTokenInfo(ctx context.Context, t string) (tokenPackage.Token, error)
	GetExternalUserID(token string) (string, error)
}

type AuthMiddleware struct {
	tokenService tokenService
	logger       logger
}

func NewAuthMiddleware(l logger, t tokenService) *AuthMiddleware {
	return &AuthMiddleware{
		logger:       l,
		tokenService: t,
	}
}

func (a *AuthMiddleware) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		tokenString, err := splitToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, invalidToken)
		}

		tokenEntity, err := a.tokenService.GetTokenInfo(c.Request().Context(), tokenString)
		if err != nil {
			if errors.Is(err, tokenPackage.ErrNoTokenFound) {
				return echo.NewHTTPError(http.StatusUnauthorized, noTokenFound)
			}

			return echo.NewHTTPError(http.StatusInternalServerError, internalServerError)
		}

		if tokenEntity.IsBlocked {
			return echo.NewHTTPError(http.StatusUnauthorized, tokenIsBlocked)
		}

		externalUserID, err := a.tokenService.GetExternalUserID(tokenString)
		if err != nil {

			if errors.Is(err, tokenPackage.ErrTokenIsInvalid) {
				return echo.NewHTTPError(http.StatusUnauthorized, invalidToken)
			}

			if errors.Is(err, tokenPackage.ErrTokenIsExpired) {
				return echo.NewHTTPError(http.StatusUnauthorized, expireToken)
			}

			return echo.NewHTTPError(http.StatusUnauthorized, invalidToken)
		}
		if tokenEntity.ExternalUserID != externalUserID {
			return echo.NewHTTPError(http.StatusUnauthorized, invalidToken)
		}

		c.Set("userID", tokenEntity.InternalUserID)
		return next(c)
	}
}

func splitToken(token string) (string, error) {
	if token == "" {
		return "", tokenPackage.ErrTokenIsInvalid
	}

	arr := strings.Split(token, " ")
	if len(arr) != 2 {
		return "", tokenPackage.ErrTokenIsInvalid
	}

	if arr[0] != "Bearer" {
		return "", tokenPackage.ErrTokenIsInvalid
	}

	if strings.Count(arr[1], ".") != 2 {
		return "", tokenPackage.ErrTokenIsInvalid
	}

	return arr[1], nil
}
