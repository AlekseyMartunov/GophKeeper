package tokenrouter

import "github.com/labstack/echo/v4"

type tokenHandlers interface {
	Lock(c echo.Context) error
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
	e.POST("token/lock", tk.handlers.Lock, tk.auth.CheckAuth, tk.log.Logging)
}
