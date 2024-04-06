package tokenrouter

import "github.com/labstack/echo/v4"

type tokenHandlers interface {
	Lock(c echo.Context) error
	CreateToken(c echo.Context) error
}

type loggerMiddleware interface {
	Logging(next echo.HandlerFunc) echo.HandlerFunc
}

type authMiddleware interface {
	CheckAuth(next echo.HandlerFunc) echo.HandlerFunc
}

type TokenControllerHTTP struct {
	handlers tokenHandlers
	auth     authMiddleware
	log      loggerMiddleware
}

func NewTokenControllerHTTP(h tokenHandlers, l loggerMiddleware, a authMiddleware) *TokenControllerHTTP {
	return &TokenControllerHTTP{
		handlers: h,
		auth:     a,
		log:      l,
	}
}

func (tk *TokenControllerHTTP) Route(e *echo.Echo) {
	e.DELETE("token", tk.handlers.Lock, tk.auth.CheckAuth, tk.log.Logging)
	e.POST("token", tk.handlers.CreateToken, tk.log.Logging)
}
