package routes

import (
	"money-tracker/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)

	app.Get("/user/:userId", controllers.GetUser)

	app.Patch("/user/:userId", controllers.UpdateUser)
}
