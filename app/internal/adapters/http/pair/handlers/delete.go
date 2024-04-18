package pairhandlers

import (
	"GophKeeper/app/internal/entity/pairs"
	"errors"
	"net/http"
)

func (ph *PairHandler) Delete(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	name := c.Param("name")

	err := ph.service.Delete(c.Request().Context(), name, userID)
	if err != nil {
		if errors.Is(err, pairs.ErrPairNothingToDelete) {
			return c.JSON(http.StatusNoContent, noContent)
		}
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	return c.JSON(http.StatusOK, messageOk)
}
