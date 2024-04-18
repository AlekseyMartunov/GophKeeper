package tokenhandlers

import (
	"GophKeeper/internal/entity/token"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (th *TokenHandler) CreateToken(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		th.log.Error(err)
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	dto := createTokenDTO{}
	if err = json.Unmarshal(b, &dto); err != nil {
		return c.JSON(http.StatusBadRequest, requestParsingError)
	}

	u, err := th.userService.GetUserInfo(c.Request().Context(), dto.Login, dto.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, unauthorized)
	}

	tokenString, err := th.tokenService.CreateAndSave(
		c.Request().Context(),
		u,
		c.Request().RemoteAddr,
		dto.TokenName,
	)

	if err != nil {
		if errors.Is(err, token.ErrTokenAlreadyExists) {
			return c.JSON(http.StatusConflict, alreadyHaveToken)
		}
		return c.JSON(http.StatusInternalServerError, internalServerError)
	}

	t := tokenDTO{
		Token: fmt.Sprintf("Bearer %s", tokenString),
	}

	return c.JSON(http.StatusOK, t)

}
