package userrouter

import (
	"github.com/labstack/echo/v4"
)

type userHandlers interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}

type UserControllerHTTP struct {
	handlers userHandlers
}

func NewUserControllerHTTP(uh userHandlers) *UserControllerHTTP {
	return &UserControllerHTTP{
		handlers: uh,
	}
}

func (uc *UserControllerHTTP) Route(e *echo.Echo) {
	e.POST("users/register", uc.handlers.Register)
	e.POST("users/login", uc.handlers.Login)

}
