package utils

import (
	"database/sql"
	"log"
	"todo/models"
)

func UserValidation(user models.User) bool {
	db, err := sql.Open("postgres", "user=postgres password=dana1234 host=localhost dbname=todos sslmode=disable")
	if err != nil {

		log.Fatalf("Error: Unable to connect to database: %v", err)

	}
	defer db.Close()

	name := ""
	if err := db.QueryRow("SELECT name FROM users WHERE email = $1", user.Email).Scan(&name); err != nil {
		if err == sql.ErrNoRows {
			return true
		}
		log.Fatalf("Error: Unable to select form database: %v", err)
	}

	return false

}
