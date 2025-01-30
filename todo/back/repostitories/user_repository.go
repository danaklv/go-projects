package repositories

import (
	"log"
	"td/back/models"

	"golang.org/x/crypto/bcrypt"
)

func InsertUserIntoDb(user models.User) error {
	err := Db.AutoMigrate(&models.User{})

	if err != nil {
		return err
	}

	hashPassword := string(HashPassword(user.Password))

	res := Db.Create(&models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: hashPassword,
	})

	if res.Error != nil {
		return res.Error
	}
	return nil

}

func HashPassword(password string) []byte {

	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal("Ошибка при хешировнании пароля", err)
	}

	return res

}
