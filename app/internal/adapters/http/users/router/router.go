package userrouter

import (
	"github.com/labstack/echo/v4"
)

type userHandlers interface {
	Register(c echo.Context) error
}

type loggerMiddleware interface {
	Logging(next echo.HandlerFunc) echo.HandlerFunc
}

type UserControllerHTTP struct {
	handlers   userHandlers
	middleware loggerMiddleware
}

func NewUserControllerHTTP(uh userHandlers, m loggerMiddleware) *UserControllerHTTP {
	return &UserControllerHTTP{
		handlers:   uh,
		middleware: m,
	}
}

func (uc *UserControllerHTTP) Route(e *echo.Echo) {
	e.POST("users/register", uc.handlers.Register, uc.middleware.Logging)
}
