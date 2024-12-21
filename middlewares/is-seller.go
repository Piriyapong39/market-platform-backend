package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func IsSeller(c *fiber.Ctx) error {
	isSeller := c.Locals("is_seller", false).(bool)
	if !isSeller {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Only seller able to access"})
	}
	return c.Next()
}
