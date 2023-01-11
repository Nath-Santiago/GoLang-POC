package users

import (
	"postgres/api/companies"
	"postgres/utils"

	db "postgres/database"

	"github.com/gofiber/fiber/v2"
)

// TODO: Use Dynamic and Raw SQL

var users []User
var user User

var company companies.Company

func GetUser(c *fiber.Ctx) error {

	params, _ := c.ParamsInt("id")
	user.ID = uint(params)

	result := db.Connect.Preload(utils.GetStructName(company)).Find(&user) // preload

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "No Data Found", "data": result})
	}
	return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Found Data", "data": user})
}

func GetAllUsers(c *fiber.Ctx) error {
	result := db.Connect.Preload(utils.GetStructName(company)).Find(&users)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something went wrong", "data": result.Error})

	}
	return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Found Data", "data": &users})
}

func CreateUser(c *fiber.Ctx) error {
	tempNewUser := new(WriteUserInputStructType)
	if err := c.BodyParser(tempNewUser); err != nil {
		panic(err)
	}
	newUser := User{
		FirstName:  tempNewUser.FirstName,
		MiddleName: tempNewUser.MiddleName,
		LastName:   tempNewUser.LastName,
		CompanyID:  uint(tempNewUser.CompanyID),
	}

	result := db.Connect.Create(&newUser)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something went wrong while creating user", "data": result.Error})

	}
	return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Created a new user", "data": newUser})
}

func UpdateUser(c *fiber.Ctx) error {

	toUpdates := new(WriteUserInputStructType)
	if err := c.BodyParser(toUpdates); err != nil {
		panic(err)
	}
	params, _ := c.ParamsInt("id")
	user.ID = uint(params)

	updatedUser := User{
		FirstName:  toUpdates.FirstName,
		MiddleName: toUpdates.MiddleName,
		LastName:   toUpdates.LastName,
		CompanyID:  toUpdates.CompanyID,
	}

	result := db.Connect.Model(user).Preload(utils.GetStructName(company)).Updates(&updatedUser)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something went wrong while updating user", "data": result.Error})

	}
	return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Updated a user", "data": updatedUser})

}

func DeleteUser(c *fiber.Ctx) error {
	params, _ := c.ParamsInt("id")
	user.ID = uint(params)

	result := db.Connect.Delete(&user) // Will not be totally deleted, 'deleted_at' column will just be updated
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something went wrong while deleting user", "data": result.Error})

	}
	return c.Status(200).JSON(fiber.Map{"status": "Success", "message": "Deleted a user", "data": user})
}
