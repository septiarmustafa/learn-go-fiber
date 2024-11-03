package main

import (
	"learn-go-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// init route
	routes.RouteInit(app)
	app.Listen(":8080")
}
