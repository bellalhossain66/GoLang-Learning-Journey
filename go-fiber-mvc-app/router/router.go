package router

import(
	"github.com/gofiber/fiber/v2"
	"go-fiber-mvc-app/handler"
)

func AppRoute(app *fiber.App) {

	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api := app.Group("/api")
	api.Post("/login", handler.LoginUser)
}