package routes

import "github.com/gofiber/fiber/v2"

func InitRoutes(app *fiber.App) {
	UserRoute(app)
	StockRoute(app)
	TransactionRoute(app)
}
