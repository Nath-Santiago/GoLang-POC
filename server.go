package main

import (
	"fmt"
	"postgres/api/authorization"
	"postgres/api/companies"
	"postgres/api/users"
	"postgres/config"

	"time"

	"github.com/gofiber/fiber/v2"
)

func StartServer() { // Echo
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
		ReadTimeout:   time.Duration(5000), // Timeout after 5 secs
	})

	users.UserRoutes(app)
	authorization.AuthRoutes(app)
	companies.CompanyRoutes(app)

	serverPort := config.GetConfig("SERVER_PORT")
	serverPortString := fmt.Sprintf(":%s", serverPort)

	app.Listen(serverPortString)
}
