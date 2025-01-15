package services

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/smtp"
	"time"
	"todo/repositories"
)

func RequestPasswordReset(email string) error {

	code, err := GenerateResetCode()
	if err != nil {
		log.Fatal(err)
	}

	expiry := time.Now().Add(15 * time.Minute)

	err = repositories.InsertInPasswordReset(email, code, expiry)

	if err != nil {
		log.Fatal(err)
	}

	return SendResetEmail(email, code)

}

func GenerateResetCode() (string, error) {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func SendResetEmail(email, code string) error {

	smtpHost := "smtp.gmail.com"

	smtpPort := "587"

	hostEmail := "kalykovadana3@gmail.com"

	password := "poxg nuft lehi qazb"

	to := []string{email}

	message := []byte("Code for Password Reset:" + "\n" + code)

	auth := smtp.PlainAuth("", hostEmail, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, hostEmail, to, message)

	return err

}

func ResetPassword(code, newPassword, email string) string {

	correct := repositories.CheckCodeInDb(code, email)

	if correct {
		repositories.ChangeUserPassword(email, newPassword)
		return "Password changed successfully"
	} else {
		return "Wrong Code"
	}

}
