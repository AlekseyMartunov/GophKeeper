package userhandlers

import (
	"GophKeeper/internal/server/entity/users"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"strings"
)

func (uh *UserHandler) Login(c echo.Context) error {
	dto := userDTO{}

	number, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	if err := json.Unmarshal(number, &dto); err != nil {
		c.JSON(http.StatusBadRequest, requestParsingError)
	}

	externalID, err := uh.service.GetExternalID(c.Request().Context(), dto.ToEntity())
	if err != nil {
		if errors.Is(err, users.ErrUserDoseNotExist) {
			return c.JSON(http.StatusNoContent, userDoseNotExist)
		}
		return c.JSON(fiber.StatusInternalServerError, internalServerError)
	}

	token, err := uh.jwt.CreateToken(externalID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, internalServerError)
	}

	return c.JSON(http.StatusOK, strings.Join([]string{"Your Token: Bearer", token}, " "))
}
