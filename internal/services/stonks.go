package services

import (
	"github.com/gofiber/fiber/v2"
)


func GetStocksHealth(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "stonks ok",
	})
}
