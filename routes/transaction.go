package routes

import (
	"money-tracker/controllers"

	"github.com/gofiber/fiber/v2"
)

func TransactionRoute(app *fiber.App) {
	app.Post("/transaction/:userId", controllers.AddTransaction)

	app.Get("/transaction/:userId", controllers.GetTransactions)

	app.Delete("/transaction/:userId/:transactionId", controllers.DeleteTransaction)

	app.Patch("/transaction/:userId/:transactionId", controllers.UpdateTransaction)

	app.Get("/transaction/:userId/:transactionId", controllers.GetTransaction)

	app.Get("/transaction/:userId/categoryTotal/:category", controllers.CalculateTotalCategory)

	app.Get("/transaction/:userId/spendFromTotal/:spendFrom", controllers.CalculateTotalSpendFrom)
}
