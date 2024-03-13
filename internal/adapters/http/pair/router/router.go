package pairrouter

import "github.com/labstack/echo/v4"

type pairHandlers interface {
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

type PairControllerHTTP struct {
	handlers pairHandlers
	auth     authMiddleware
	logger   loggerMiddleware
}

func NewPairControllerHTTP(h pairHandlers, l loggerMiddleware, a authMiddleware) *PairControllerHTTP {
	return &PairControllerHTTP{
		handlers: h,
		auth:     a,
		logger:   l,
	}
}

func (pc *PairControllerHTTP) Route(e *echo.Echo) {
	e.POST("pairs", pc.handlers.Save, pc.auth.CheckAuth, pc.logger.Logging)
	e.GET("pairs/:name", pc.handlers.Get, pc.auth.CheckAuth, pc.logger.Logging)
	e.GET("pairs", pc.handlers.GetAll, pc.auth.CheckAuth, pc.logger.Logging)
	e.DELETE("pairs/:name", pc.handlers.Delete, pc.auth.CheckAuth, pc.logger.Logging)
}
