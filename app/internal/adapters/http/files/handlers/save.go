package filehandlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (fh *FileHandler) Save(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		fh.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	dto := fileDTO{}
	dto.userID = userID

	if err = json.Unmarshal(b, &dto); err != nil {
		fh.log.Error(err)
		return c.JSON(http.StatusBadRequest, requestParsingError)
	}

	err = fh.service.Save(c.Request().Context(), dto.ToEntity())
	if err != nil {
		fh.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}
	return nil
}
