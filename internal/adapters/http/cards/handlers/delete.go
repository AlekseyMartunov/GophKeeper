package cardhandlers

import (
	"GophKeeper/internal/entity/card"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ch *CardHandler) Delete(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	name := c.Param("name")

	err := ch.service.Delete(c.Request().Context(), name, userID)
	if err != nil {
		if errors.Is(err, card.ErrCardNothingToDelete) {
			return c.JSON(http.StatusNoContent, card.ErrCardNothingToDelete)
		}
		ch.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}
	return c.JSON(http.StatusOK, messageOk)

}
