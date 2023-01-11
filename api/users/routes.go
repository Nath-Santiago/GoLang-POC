package users

import (
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(f *fiber.App) {
	// TODO: Group This... Issue, bug when using group, laging base route natatawag.
	// userRoute := f.Group("/users")
	// userRoute.Post("/create", CreateUser)
	// userRoute.Get("/:id", GetUser)
	// userRoute.Put("/update/:id", UpdateUser)
	// userRoute.Delete("/delete/:id", DeleteUser)

	f.Get("/users", GetAllUsers)
	/*C*/ f.Post("/users/create", CreateUser) // TODO?: Batch Insert?
	/*R*/ f.Get("/users/:id", GetUser)
	/*U*/ f.Put("users/update/:id", UpdateUser)
	/*D*/ f.Delete("users/delete/:id", DeleteUser)
}
