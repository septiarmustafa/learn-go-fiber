package routes

import (
	"learn-go-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", handler.UserHandlerRead)
}
