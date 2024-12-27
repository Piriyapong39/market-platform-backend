package product

import (

	// import package from external
	"github.com/gofiber/fiber/v2"

	// import middlewares
	"github.com/piriyapong39/market-platform/middlewares"
)

func ProductRoute(app *fiber.App) {
	product := app.Group("/product")

	product.Use(middlewares.Authentication)
	product.Use(middlewares.IsSeller)
	product.Post("/create-product", createProduct)
}
