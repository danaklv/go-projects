package repositories


import (
	"database/sql"
	"fmt"
	"log"
	"todo/models"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func ConnectToDb() {
	db, err := sql.Open("postgres", "user=postgres password=dana1234 host=localhost dbname=todos sslmode=disable")
	if err != nil {
		fmt.Println("here")
		log.Fatalf("Error: Unable to connect to database: %v", err)
	}
	defer db.Close()

	res, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY,name VARCHAR  NOT NULL,email VARCHAR  NOT NULL,password VARCHAR  NOT NULL)")

	if err != nil {
		log.Fatalf("Error: Unable to create table: %v", err)
	}

	res.RowsAffected()

}

func InsertUserIntoDb(user *models.User) {
	db, err := sql.Open("postgres", "user=postgres password=dana1234 host=localhost dbname=todos sslmode=disable")
	if err != nil {
		fmt.Println("im here")
		log.Fatalf("Error: Unable to connect to database: %v", err)
	}
	defer db.Close()

	hashedPassword := string(HashPassword(user.Password))

	stmt := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`
	res, err := db.Exec(stmt, user.Name, user.Email, hashedPassword)
	if err != nil {
		log.Fatalf("Error: Unable to INSERT INTO users: %v", err)
	}

	res.RowsAffected()
}

func ChechUserInDb(email, password string) bool {
	db, err := sql.Open("postgres", "user=postgres password=dana1234 host=localhost dbname=todos sslmode=disable")
	if err != nil {
		fmt.Println("im here")
		log.Fatalf("Error: Unable to connect to database: %v", err)
	}
	defer db.Close()

	passwordInDb := ""
	if err := db.QueryRow("SELECT password FROM users WHERE email = $1", email).Scan(&passwordInDb); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Fatalf("Error: Unable to select form database: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordInDb), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func HashPassword(password string) []byte {

	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal("Ошибка при хешировнании пароля", err)
	}

	return res

}

func SelectNameFromDb(email string) string {
	db, err := sql.Open("postgres", "user=postgres password=dana1234 host=localhost dbname=todos sslmode=disable")
	if err != nil {
		
		log.Fatalf("Error: Unable to connect to database: %v", err)
	}
	defer db.Close()

	name := ""
	if err := db.QueryRow("SELECT name FROM users WHERE email = $1", email).Scan(&name); err != nil {
		if err == sql.ErrNoRows {
			return ""
		}
		log.Fatalf("Error: Unable to select form database: %v", err)
	}
	return name

}
