package companies

import (
	db "postgres/database"

	"github.com/gofiber/fiber/v2"
)

var companies []Company

// var company Company

func CreateCompany(c *fiber.Ctx) error {
	input := new(CreateCompanyInput)
	if err := c.BodyParser(input); err != nil { // insert validation here
		panic(err)
	}
	newCompany := Company{
		Name: input.Name,
	}

	result := db.Connect.Create(&newCompany)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something went wrong while creating company", "data": result.Error})

	}
	return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Created a new company", "data": newCompany})
}

func GetAllCompany(c *fiber.Ctx) error {
	result := db.Connect.Find(&companies)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something went wrong", "data": result.Error})

	}
	return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Found Data", "data": companies})
}
