package main

import (
	"learn-go-fiber/database"
	"learn-go-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// init database
	database.Database()

	app := fiber.New()

	// init route
	routes.RouteInit(app)
	app.Listen(":8080")
}
