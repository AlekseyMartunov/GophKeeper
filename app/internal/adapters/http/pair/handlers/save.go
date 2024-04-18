package pairhandlers

import (
	"GophKeeper/app/internal/entity/pairs"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (ph *PairHandler) Save(c echo.Context) error {
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

	if err = json.Unmarshal(b, &dto); err != nil {
		return c.JSON(http.StatusBadRequest, requestParsingError)
	}
	dto.UserID = userID

	err = ph.service.Save(c.Request().Context(), dto.toEntity())
	if err != nil {
		if errors.Is(err, pairs.ErrPairAlreadyExists) {
			return c.JSON(http.StatusConflict, pairAlreadyExists)
		}
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}
	return c.JSON(http.StatusOK, messageOk)
}
