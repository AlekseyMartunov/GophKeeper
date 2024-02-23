package userhandlers

import "github.com/gofiber/fiber/v2"

func (uh *UserHandler) Register(c *fiber.Ctx) error {
	dto := userDTO{}

	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	if err := uh.service.Save(c.Context(), dto.ToEntity()); err != nil {
		return err
	}

	return c.JSON(dto)
}
