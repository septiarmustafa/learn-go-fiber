package main

import (
	"learn-go-fiber/database"
	"learn-go-fiber/database/migration"
	"learn-go-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// init database
	database.Database()
	migration.RunMigration()

	app := fiber.New()

	// init route
	routes.RouteInit(app)
	app.Listen(":8080")
}
