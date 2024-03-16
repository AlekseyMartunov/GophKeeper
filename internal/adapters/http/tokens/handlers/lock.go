package tokenhandlers

import (
	"GophKeeper/internal/entity/token"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (th *TokenHandler) Lock(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		th.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	dto := dtoToken{}

	err = json.Unmarshal(b, &dto)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestParsingError)
	}

	err = th.service.LockToken(c.Request().Context(), dto.Token, dto.Status, userID)
	if err != nil {
		if errors.Is(err, token.ErrNoTokenFound) {
			return c.JSON(http.StatusNoContent, noToken)
		}
		th.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}
	return c.JSON(http.StatusOK, messageOk)
}
