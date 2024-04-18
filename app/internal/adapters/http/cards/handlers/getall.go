package cardhandlers

import (
	"GophKeeper/app/internal/entity/card"
	"errors"
	"net/http"
)

func (ch *CardHandler) GetAll(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	cardArr, err := ch.service.GetAll(c.Request().Context(), userID)
	if err != nil {
		if errors.Is(err, card.ErrCardDoseNotExist) {
			return c.JSON(http.StatusNoContent, allCardsDoseNotExist)
		}
		ch.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	return c.JSON(http.StatusOK, arrDTO(cardArr))
}
