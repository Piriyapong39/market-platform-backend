package user

import (
	"github.com/gofiber/fiber/v2"
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
