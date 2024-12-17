package main

import (
	"fmt"
	"os"

	// import package from external

	// import package from internal
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	// import .env file
	err := godotenv.Load("./config/.env")
	if err != nil {
		fmt.Println(err)
	}
	port := os.Getenv("PORT")

	// import routes
	app.Get("/", greedUser)
	if err := app.Listen(":" + port); err != nil {
		fmt.Println(err)
	}
}

func greedUser(c *fiber.Ctx) error {
	return c.SendString("Hello Golang")
}
