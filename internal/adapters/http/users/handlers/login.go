package userhandlers

import (
	"GophKeeper/internal/entity/users"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
)

func (uh *UserHandler) Login(c echo.Context) error {
	dto := userDTO{}

	number, err := io.ReadAll(c.Request().Body)
	if err != nil {
		uh.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	if err := json.Unmarshal(number, &dto); err != nil {
		return c.JSON(http.StatusBadRequest, requestParsingError)
	}

	externalID, err := uh.service.GetExternalID(c.Request().Context(), dto.ToEntity())
	if err != nil {
		if errors.Is(err, users.ErrUserDoseNotExist) {
			return c.JSON(http.StatusNoContent, userDoseNotExist)
		}
		uh.log.Error(err)
		return c.JSON(fiber.StatusInternalServerError, internalServerError)
	}

	token, err := uh.jwt.CreateToken(externalID)
	if err != nil {
		uh.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	return c.JSON(http.StatusOK, strings.Join([]string{"Your Token: Bearer", token}, " "))
}