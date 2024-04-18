package pairhandlers

import (
	"GophKeeper/internal/entity/pairs"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ph *PairHandler) GetAll(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	pairsArr, err := ph.service.GetAll(c.Request().Context(), userID)
	if err != nil {
		if errors.Is(err, pairs.ErrPairDoseNotExist) {
			return c.JSON(http.StatusNoContent, allPairDoseNotExist)
		}
		ph.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	return c.JSON(http.StatusOK, arrDTO(pairsArr))
}
