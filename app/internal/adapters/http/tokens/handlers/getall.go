package tokenhandlers

import (
	"GophKeeper/internal/entity/token"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (th *TokenHandler) GetAll(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	tokens, err := th.tokenService.GetAll(c.Request().Context(), userID)
	if err != nil {
		if errors.Is(err, token.ErrNoTokenFound) {
			return c.JSON(http.StatusNoContent, noClients)
		}
		th.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	clients := clientsDTO{Clients: tokens}
	fmt.Println(tokens)

	return c.JSON(http.StatusOK, clients)
}
