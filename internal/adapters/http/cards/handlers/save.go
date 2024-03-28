package cardhandlers

import (
	"GophKeeper/internal/entity/card"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ch *CardHandler) Save(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		ch.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	dto := cardDTO{}

	if err := json.Unmarshal(b, &dto); err != nil {
		ch.log.Error(err)
		return c.JSON(http.StatusBadRequest, requestParsingError)
	}
	dto.userID = userID

	err = ch.service.Save(c.Request().Context(), dto.toEntity())
	if err != nil {
		if errors.Is(err, card.ErrCardAlreadyExists) {
			return c.JSON(http.StatusConflict, cardAlreadyExists)
		}
		ch.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}
	return c.JSON(http.StatusOK, messageOk)

}
