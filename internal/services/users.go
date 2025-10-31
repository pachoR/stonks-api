package services

import (
	"github.com/gofiber/fiber/v2"
)

func GetUserHealth(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "users ok",
	})
}
