package routes

import (
	"money-tracker/controllers"
	"money-tracker/middleware"

	"github.com/gofiber/fiber/v2"
)

func TransactionRoute(router fiber.Router) {
	router.Post("/:userId", middleware.AuthenticateUser,controllers.AddTransaction)

	router.Get("/:userId", middleware.AuthenticateUser,controllers.GetTransactions)

	router.Delete("/:userId/:transactionId", middleware.AuthenticateUser,controllers.DeleteTransaction)

	router.Patch("/:userId/:transactionId", middleware.AuthenticateUser,controllers.UpdateTransaction)

	router.Get("/:userId/:transactionId", middleware.AuthenticateUser,controllers.GetTransaction)

	router.Get("/:userId/categoryTotal/:category", middleware.AuthenticateUser,controllers.CalculateTotalCategory)

	router.Get("/:userId/spendFromTotal/:spendFrom", middleware.AuthenticateUser,controllers.CalculateTotalSpendFrom)
}
