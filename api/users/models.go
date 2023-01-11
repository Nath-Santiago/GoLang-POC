package users

import (
	"postgres/api/companies"

	"gorm.io/gorm"
)

type User struct {
	FirstName  string
	MiddleName string
	LastName   string
	gorm.Model
	CompanyID uint
	Company   companies.Company
}

type WriteUserInputStructType struct { // Editable Fields only
	FirstName  string
	MiddleName string
	LastName   string
	CompanyID  uint
}
