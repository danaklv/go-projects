package models

type User struct {
	Id       int `gorm:"primaryKey"`
	Username string
	Email    string
	Password string
}
