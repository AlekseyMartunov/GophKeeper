package router

import "github.com/labstack/echo/v4"

type fileHandlers interface {
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

type FileControllerHTTP struct {
	handlers fileHandlers
	auth     authMiddleware
	logger   loggerMiddleware
}

func NewFileControllerHTTP(h fileHandlers, l loggerMiddleware, a authMiddleware) *FileControllerHTTP {
	return &FileControllerHTTP{
		handlers: h,
		auth:     a,
		logger:   l,
	}
}

func (fc *FileControllerHTTP) Route(e *echo.Echo) {
	e.POST("file", fc.handlers.Save, fc.auth.CheckAuth, fc.logger.Logging)
	e.GET("file/:name", fc.handlers.Get, fc.auth.CheckAuth, fc.logger.Logging)
	e.GET("file", fc.handlers.GetAll, fc.auth.CheckAuth, fc.logger.Logging)
	e.DELETE("file/:name", fc.handlers.Delete, fc.auth.CheckAuth, fc.logger.Logging)
}
