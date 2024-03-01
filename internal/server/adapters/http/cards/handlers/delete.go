package cardhandlers

import (
	"GophKeeper/internal/server/entity/card"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (ch *CardHandler) Delete(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		ch.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	var name cardName

	if err := json.Unmarshal(b, &name); err != nil {
		return c.JSON(http.StatusBadRequest, requestParsingError)
	}

	err = ch.service.Delete(c.Request().Context(), name.Name, userID)
	if err != nil {
		if errors.Is(err, card.ErrCardNothingToDelete) {
			return c.JSON(http.StatusNoContent, card.ErrCardNothingToDelete)
		}
		ch.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

}
