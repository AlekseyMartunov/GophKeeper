package pairhandlers

import (
	"GophKeeper/internal/entity/pairs"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (ph *PairHandler) Delete(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	b, err := io.ReadAll(c.Request().Body)

	if err != nil {
		ph.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	var dto nameDTO

	if err := json.Unmarshal(b, &dto); err != nil {
		return c.JSON(http.StatusBadRequest, requestParsingError)
	}

	err = ph.service.Delete(c.Request().Context(), dto.Name, userID)
	if err != nil {
		if errors.Is(err, pairs.ErrPairNothingToDelete) {
			return c.JSON(http.StatusNoContent, noContent)
		}
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	return c.JSON(http.StatusOK, messageOk)
}
