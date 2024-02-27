package pairhandlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"GophKeeper/internal/server/entity/pairs"

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

	dto := pairDTO{}

	if err := json.Unmarshal(b, &dto); err != nil {
		return c.JSON(http.StatusBadRequest, requestParsingError)
	}
	dto.UserID = userID

	pair, err := ph.service.Get(c.Request().Context(), dto.Name, dto.UserID)
	if err != nil {
		if errors.Is(err, pairs.ErrPairDoseNotExist) {
			return c.JSON(http.StatusNoContent, pairDoseNotExist)
		}
		ph.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	dto.fromEntity(pair)
	return c.JSON(http.StatusOK, dto)
}
