package userhandlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"GophKeeper/internal/server/entity/users"

	"github.com/labstack/echo/v4"
)

func (uh *UserHandler) Register(c echo.Context) error {
	dto := userDTO{}

	number, err := io.ReadAll(c.Request().Body)
	if err != nil {
		uh.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	if err := json.Unmarshal(number, &dto); err != nil {
		return c.JSON(http.StatusBadRequest, requestParsingError)
	}

	if err := uh.service.Save(c.Request().Context(), dto.ToEntity()); err != nil {
		if errors.Is(err, users.ErrUserAlreadyExists) {
			return c.JSON(http.StatusConflict, userAlreadyExists)
		}
		uh.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	return c.JSON(http.StatusOK, dto)
}
