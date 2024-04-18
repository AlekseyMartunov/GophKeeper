package pairhandlers

import (
	"GophKeeper/internal/entity/pairs"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ph *PairHandler) Get(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	name := c.Param("name")
	pair, err := ph.service.Get(c.Request().Context(), name, userID)
	if err != nil {
		if errors.Is(err, pairs.ErrPairDoseNotExist) {
			return c.JSON(http.StatusNoContent, pairDoseNotExist)
		}
		ph.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}
	dto := pairDTO{}

	dto.fromEntity(pair)
	return c.JSON(http.StatusOK, dto)
}
