package authorization

import (
	"postgres/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(f *fiber.App) {

	// TODO: Group This
	f.Get("/auth", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	f.Post("/auth/register", RegisterUser)
	f.Get("/auth/:id", GetUser)

	f.Post("/auth/login", LoginUser)

	f.Patch("/auth/:id", middlewares.Protected(), UpdateUser)
	f.Delete("/auth/:id", middlewares.Protected(), DeleteUser)

}
