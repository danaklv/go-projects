package repositories

import (
	"database/sql"
	"log"
	"time"
	"todo/models"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func ConnectToDb() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=dana1234 host=localhost dbname=todos sslmode=disable")
	if err != nil {
		log.Fatalf("Error: Unable to connect to database: %v", err)
	}

}

func InsertUserIntoDb(user *models.User) {

	hashedPassword := string(HashPassword(user.Password))

	stmt := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`
	_, err := DB.Exec(stmt, user.Name, user.Email, hashedPassword)
	if err != nil {
		log.Fatalf("Error: Unable to INSERT INTO users: %v", err)
	}

}

func CheckUserInDb(email, password string) bool {

	passwordInDb := ""
	if err := DB.QueryRow("SELECT password FROM users WHERE email = $1", email).Scan(&passwordInDb); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Fatalf("Error: Unable to select form database: %v", err)
	}

	err := bcrypt.CompareHashAndPassword([]byte(passwordInDb), []byte(password))

	return err == nil
}

func HashPassword(password string) []byte {

	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal("Ошибка при хешировнании пароля", err)
	}

	return res

}

func GetUserIdFromDb(email string) (int, error) {
	id := 0
	if err := DB.QueryRow("SELECT id FROM users WHERE email = $1", email).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		}
		log.Fatalf("Error: Unable to select form database: %v", err)
	}
	return id, nil

}

func InsertInPasswordReset(email, code string, expiry time.Time) error {
	stmt := `INSERT INTO password_reset (email, code, expires_at) VALUES ($1,$2,$3)`
	_, err := DB.Exec(stmt, email, code, expiry)
	return err

}

func CheckCodeInDb(code, email string) bool {

	emailBd := ""
	query := `SELECT email FROM password_reset WHERE code = $1`

	err := DB.QueryRow(query, code).Scan(&emailBd)

	if err != nil {
		log.Fatal("Error select from password_reset: ", err)
	}

	if emailBd != email {
		return false
	}
	return true

}

func ChangeUserPassword(email, password string )  {

	hashPassword := HashPassword(password)

	stmt := `UPDATE users SET password = $1 WHERE email = $2`

	_, err := DB.Exec(stmt, hashPassword, email)

	if err != nil {
		log.Fatal(err)
	}

	

}
