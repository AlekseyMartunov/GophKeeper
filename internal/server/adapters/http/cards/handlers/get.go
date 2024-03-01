package cardhandlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"GophKeeper/internal/server/entity/card"
)

func (ch *CardHandler) Get(c echo.Context) error {
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

	res, err := ch.service.Get(c.Request().Context(), name.Name, userID)
	if err != nil {
		if errors.Is(err, card.ErrCardDoseNotExist) {
			return c.JSON(http.StatusNoContent, cardDoseNotExist)
		}
		ch.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	dto := cardDTO{}
	dto.fromEntity(res)

	return c.JSON(http.StatusOK, dto)
}
