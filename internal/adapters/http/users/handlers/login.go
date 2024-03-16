package userhandlers

import (
	"GophKeeper/internal/entity/token"
	"GophKeeper/internal/entity/users"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (uh *UserHandler) Login(c echo.Context) error {
	dto := userDTO{}

	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		uh.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	if err = json.Unmarshal(b, &dto); err != nil {
		return c.JSON(http.StatusBadRequest, requestParsingError)
	}

	user, err := uh.service.GetUserInfo(c.Request().Context(), dto.UserLogin, dto.UserPassword)
	if err != nil {
		if errors.Is(err, users.ErrUserDoseNotExist) {
			return c.JSON(http.StatusNoContent, userDoseNotExist)
		}
		uh.log.Error(err)
		return c.JSON(fiber.StatusInternalServerError, internalServerError)
	}

	tokenString, err := uh.tokenService.CreateAndSave(c.Request().Context(), user, dto.NameForToken)
	if err != nil {
		if errors.Is(err, token.ErrTokenAlreadyExists) {
			return c.JSON(http.StatusConflict, loginAlreadyExists)
		}
		uh.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	t := tokenDTO{
		Token: fmt.Sprintf("Bearer %s", tokenString),
	}

	return c.JSON(http.StatusOK, t)
}
