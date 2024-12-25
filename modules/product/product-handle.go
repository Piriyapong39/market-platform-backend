package product

import (
	"github.com/gofiber/fiber/v2"

	// import service

	userservices "github.com/piriyapong39/market-platform/services/user-services"
)

func createProduct(c *fiber.Ctx) error {
	productRequest := new(Product)
	token := c.Get("Authorization")
	userData, err := userservices.VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	productRequest.User_id = userData.Id
	if err := c.BodyParser(&productRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if productRequest.Name == "" || productRequest.Description == "" || productRequest.Stock == 0 || productRequest.Price == 0 || productRequest.CategoryID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required fields",
		})
	}
	if len(productRequest.PicPath) > 5 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Maximum 5 images are allowed",
		})
	}
	result, err := _createProduct(*productRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Product created successfully",
		"data":    result,
	})
}
