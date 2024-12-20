package user

import (
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	user := app.Group("/user")
	user.Post("/register", userRegister)
	user.Post("/login", userLogin)

}
