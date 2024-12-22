package middlewares

import (
	"github.com/gofiber/fiber/v2"

	userservices "github.com/piriyapong39/market-platform/services/user-services"
)

func Authentication(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	userData, err := userservices.VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	c.Locals("is_seller", userData.Is_seller)
	return c.Next()
}
