package cardhandlers

import (
	"GophKeeper/internal/entity/card"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ch *CardHandler) Get(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	name := c.Param("name")
	res, err := ch.service.Get(c.Request().Context(), name, userID)
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
