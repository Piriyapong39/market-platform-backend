package main

import (

	// import package from internal
	"fmt"
	"os"

	// import package from external
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	// import route
	"github.com/piriyapong39/market-platform/modules/product"
	"github.com/piriyapong39/market-platform/modules/user"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	// import .env file
	err := godotenv.Load("./config/.env")
	if err != nil {
		fmt.Println(err)
	}
	port := os.Getenv("PORT")

	// active routes
	user.UserRoute(app)
	product.ProductRoute(app)

	// start server on port
	if err := app.Listen(":" + port); err != nil {
		fmt.Println(err)
	}
}
