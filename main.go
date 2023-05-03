package main

import (
	"money-tracker/configs"
	"money-tracker/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.UserRoute(app)
	routes.StockRoute(app)

	app.Listen(":6969")
}
