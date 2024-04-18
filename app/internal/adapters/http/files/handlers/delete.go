package filehandlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (fh *FileHandler) Delete(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	fileName := c.Param("name")

	err := fh.service.Delete(c.Request().Context(), userID, fileName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	return c.JSON(http.StatusOK, messageOK)
}
