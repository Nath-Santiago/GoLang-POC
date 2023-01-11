package main

import (
	"fmt"
	"postgres/api/authorization"
	"postgres/api/companies"
	"postgres/api/users"
	"postgres/database"

	_ "github.com/lib/pq"
)

func getModels() []interface{} {
	var modelsSlice []interface{}

	// Add Models here and append to slice; For Auto Migrating
	userTable := users.User{}
	authorizedUserTable := authorization.AuthorizedUser{}
	companyTable := companies.Company{}
	// sampleTable := Sample{}

	modelsSlice = append(modelsSlice, userTable) // append sampleTable
	modelsSlice = append(modelsSlice, authorizedUserTable)
	modelsSlice = append(modelsSlice, companyTable)
	return modelsSlice
}

func migrateTables() {
	tables := getModels()
	for _, table := range tables {
		database.Connect.AutoMigrate(table)
	}

	fmt.Println("Automigration completed.")
}

func init() {
	/** Starts the Databse connection */
	database.ConnectToDatabase()
	migrateTables()
}

func main() {
	/** Starts the web app server*/
	StartServer()
}
