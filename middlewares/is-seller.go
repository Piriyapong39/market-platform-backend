package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func IsSeller(c *fiber.Ctx) error {
	isSeller := c.Locals("is_seller").(bool)
	if !isSeller {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"is_seller": false})
	}
	return c.Next()
}
