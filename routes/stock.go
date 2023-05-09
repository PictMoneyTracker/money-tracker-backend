package routes

import (
	"money-tracker/controllers"
	"money-tracker/middleware"

	"github.com/gofiber/fiber/v2"
)

func StockRoute(router fiber.Router) {
	router.Post("/:userId", middleware.AuthenticateUser, controllers.AddStock)

	router.Get("/:userId", middleware.AuthenticateUser, controllers.GetStocks)

	router.Delete("/:userId/:stockId", middleware.AuthenticateUser, controllers.DeleteStock)
}
