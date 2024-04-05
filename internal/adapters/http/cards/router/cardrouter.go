package cardrouter

import "github.com/labstack/echo/v4"

type cardHandlers interface {
	Get(c echo.Context) error
	Save(c echo.Context) error
	GetAll(c echo.Context) error
	Delete(c echo.Context) error
}

type loggerMiddleware interface {
	Logging(next echo.HandlerFunc) echo.HandlerFunc
}

type authMiddleware interface {
	CheckAuth(next echo.HandlerFunc) echo.HandlerFunc
}

type CardControllerHTTP struct {
	handlers cardHandlers
	auth     authMiddleware
	logger   loggerMiddleware
}

func NewCardControllerHTTP(h cardHandlers, l loggerMiddleware, a authMiddleware) *CardControllerHTTP {
	return &CardControllerHTTP{
		handlers: h,
		auth:     a,
		logger:   l,
	}
}

func (cc *CardControllerHTTP) Route(e *echo.Echo) {
	e.POST("card", cc.handlers.Save, cc.auth.CheckAuth, cc.logger.Logging)
	e.GET("card/:name", cc.handlers.Get, cc.auth.CheckAuth, cc.logger.Logging)
	e.GET("card", cc.handlers.GetAll, cc.auth.CheckAuth, cc.logger.Logging)
	e.DELETE("card/:name", cc.handlers.Delete, cc.auth.CheckAuth, cc.logger.Logging)
}
