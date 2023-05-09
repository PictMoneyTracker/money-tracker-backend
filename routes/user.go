package routes

import (
	"money-tracker/controllers"
	"money-tracker/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)
	app.Get("/login/:provider", controllers.HandleLogin)
	app.Get("/auth/callback/:provider", controllers.HandleAuth)

	app.Get("/user/:userId", middleware.AuthenticateUser, controllers.GetUser)

	app.Patch("/user/:userId", middleware.AuthenticateUser, controllers.UpdateUser)

}
