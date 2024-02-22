package handlers

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func (uh *UserHandler) Login(c *fiber.Ctx) error {
	dto := userDTO{}

	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	externalID, err := uh.service.GetExternalID(c.Context(), dto.ToEntity())
	if err != nil {
		return err
	}

	token, err := uh.jwt.CreateToken(externalID)
	if err != nil {
		return err
	}
	c.Set("Authorization", strings.Join([]string{"Bearer", token}, " "))

	return c.JSON(fiber.StatusOK)
}
