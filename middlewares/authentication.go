package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	userservices "github.com/piriyapong39/market-platform/services/user-services"
)

func Authentication(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
	}
	tokenPart := strings.Split(token, " ")
	if tokenPart[0] != "Bearer" && len(tokenPart) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "wrong format token"})
	}
	userData, err := userservices.VerifyToken(tokenPart[1])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	c.Locals("is_seller", userData.Is_seller)
	return c.Next()
}
