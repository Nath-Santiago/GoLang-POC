package database

import (
	"fmt"
	"postgres/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Connect *gorm.DB

func ConnectToDatabase() {

	db_user := config.GetConfig("POSTGRES_USERNAME")
	db_pass := config.GetConfig("POSTGRES_PASSWORD")
	db_name := config.GetConfig("POSTGRES_DBNAME")
	db_port := config.GetConfig("POSTGRES_PORT")
	db_host := config.GetConfig("POSTGRES_HOST")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s", db_user, db_pass, db_name, db_port, db_host)

	/** Start Postgres Connection */
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: false, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   true,  // skip the snake_casing of names
			// NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	Connect = db
}
