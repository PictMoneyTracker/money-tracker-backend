package routes

import (
	"money-tracker/controllers"

	"github.com/gofiber/fiber/v2"
)

func TransactionRoute(app *fiber.App) {
	app.Post("/transaction/:userId", controllers.AddTransaction)
}
