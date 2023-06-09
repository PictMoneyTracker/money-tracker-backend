package main

import (
	"money-tracker/configs"
	"money-tracker/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.SetUpAuth()
	configs.ConnectDB()

	routes.InitRoutes(app)

	app.Listen(":6969")
}
