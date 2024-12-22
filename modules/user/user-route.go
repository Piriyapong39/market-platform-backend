package user

import (
	"github.com/gofiber/fiber/v2"

	"github.com/piriyapong39/market-platform/middlewares"
)

func UserRoute(app *fiber.App) {
	user := app.Group("/user")
	user.Post("/register", userRegister)
	user.Post("/login", userLogin)
	user.Post("/authen", userAuthen)

	seller := app.Group("/seller")
	seller.Use(middlewares.Authentication)
	seller.Post("/confirm-to-seller", ConfirmToSeller)
	seller.Use(middlewares.IsSeller)
	seller.Post("/authen", sellerAuthen)
}
