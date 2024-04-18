package filehandlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (fh *FileHandler) GetAll(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	files, err := fh.service.GetAll(c.Request().Context(), userID)
	if err != nil {
		fh.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	var arrDTO []fileDTO
	for _, f := range files {
		d := fileDTO{
			Name:        f.Name,
			CreatedTime: f.CreatedTime,
		}
		arrDTO = append(arrDTO, d)
	}

	return c.JSON(http.StatusOK, arrDTO)

}
