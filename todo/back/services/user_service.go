package services

import (
	"log"
	"os"
	"td/back/models"
	repositories "td/back/repostitories"
)

func RegisterUserService(user models.User) {

	err := repositories.InsertUserIntoDb(user)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
