package migration

import (
	"learn-go-fiber/database"
	"learn-go-fiber/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)
	}
	log.Println("Database migrated")
}
