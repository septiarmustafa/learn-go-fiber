package handler

import (
	"learn-go-fiber/database"
	"learn-go-fiber/model/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(ctx *fiber.Ctx) error {
	var users []entity.User
	err := database.DB.Find(&users).Error
	if err != nil {
		log.Println(err)
	}
	return ctx.JSON(users)
}
