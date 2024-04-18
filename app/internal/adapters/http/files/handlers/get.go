package filehandlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (fh *FileHandler) Get(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	fileName := c.Param("name")

	file, err := fh.service.Get(c.Request().Context(), userID, fileName)
	if err != nil {
		fh.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	dto := fileDTO{}
	dto.FromEntity(*file)

	return c.JSON(http.StatusOK, dto)
}
