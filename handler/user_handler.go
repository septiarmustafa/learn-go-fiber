package handler

import (
	"learn-go-fiber/database"
	"learn-go-fiber/model/entity"
	"learn-go-fiber/model/request"
	"log"

	"github.com/go-playground/validator/v10"
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

func CreateUser(ctx *fiber.Ctx) error {
	user := new(request.UserRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(
			fiber.Map{
				"message": "failed",
				"error":   errValidate.Error(),
			},
		)
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	err := database.DB.Create(&newUser).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}
