package routes

import (
	"money-tracker/controllers"

	"github.com/gofiber/fiber/v2"
)

func StockRoute(app *fiber.App) {
	app.Post("/stock/:userId", controllers.AddStock)

	app.Get("/stock/:userId", controllers.GetStocks)

	app.Delete("/stock/:userId/:stockId", controllers.DeleteStock)
}
