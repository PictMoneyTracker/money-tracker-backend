package routes

import "github.com/gofiber/fiber/v2"

func UserRoute(app *fiber.App) {
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}
