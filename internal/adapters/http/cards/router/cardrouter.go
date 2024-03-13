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
	e.POST("card/save", cc.handlers.Save, cc.logger.Logging, cc.auth.CheckAuth)
	e.POST("card/get", cc.handlers.Get, cc.logger.Logging, cc.auth.CheckAuth)
	e.POST("card/getall", cc.handlers.GetAll, cc.logger.Logging, cc.auth.CheckAuth)
	e.DELETE("card/delete", cc.handlers.Delete, cc.logger.Logging, cc.auth.CheckAuth)
}
