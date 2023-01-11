package companies

import (
	"github.com/gofiber/fiber/v2"
)

func CompanyRoutes(f *fiber.App) {

	f.Get("/company", GetAllCompany)
	/*C*/
	f.Post("/company/create", CreateCompany)
	// /*R*/ f.Get("/company/:id", GetCompany)
	// /*U*/ f.Put("company/update/:id", UpdateCompany)
	// /*D*/ f.Delete("company/delete/:id", DeleteCompany)
}
