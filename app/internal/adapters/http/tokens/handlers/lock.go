package tokenhandlers

import (
	"GophKeeper/app/internal/entity/token"
	"encoding/json"
	"errors"
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

	dto := tokenDeleteDTO{}

	err = json.Unmarshal(b, &dto)
	if err != nil {
		return c.JSON(http.StatusBadRequest, requestParsingError)
	}

	err = th.tokenService.LockToken(c.Request().Context(), dto.TokenName, userID)
	if err != nil {
		if errors.Is(err, token.ErrNoTokenFound) {
			return c.JSON(http.StatusNoContent, noToken)
		}
		th.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}
	return c.JSON(http.StatusOK, messageOk)
}
