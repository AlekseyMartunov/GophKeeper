package pairhandlers

import (
	"GophKeeper/internal/entity/pairs"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ph *PairHandler) Get(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		ph.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	name := nameDTO{}

	if err := json.Unmarshal(b, &name); err != nil {
		return c.JSON(http.StatusBadRequest, requestParsingError)
	}

	pair, err := ph.service.Get(c.Request().Context(), name.Name, userID)
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
