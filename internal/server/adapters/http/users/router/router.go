package router

import (
	"github.com/gofiber/fiber/v2"
)

type userHandlers interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}

type UserControllerHTTP struct {
	handlers userHandlers
}

func NewUserControllerHTTP(uh userHandlers) *UserControllerHTTP {
	return &UserControllerHTTP{
		handlers: uh,
	}
}

func (uc *UserControllerHTTP) Route() *fiber.App {
	app := fiber.New()

	app.Post("/register", uc.handlers.Register)
	app.Post("/login", uc.handlers.Login)

	return app
}
