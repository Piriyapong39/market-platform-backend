package user

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	userservices "github.com/piriyapong39/market-platform/services/user-services"
)

func userRegister(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if user.Email == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You are missing some fields please check again"})
	}
	results, err := _userRegister(*user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"msg": results})
}

func userLogin(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "You are missing some fields please check again"})
	}

	token, err := _userLogin(*user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"token": token})
}

func userAuthen(c *fiber.Ctx) error {
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
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg":  "Authorized",
		"user": userData,
	})
}

func sellerAuthen(c *fiber.Ctx) error {
	return c.SendString("Ok")
}
