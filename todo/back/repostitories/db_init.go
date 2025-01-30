package repositories

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func ConnectToDb() error {
	dsn := "host = localhost user=postgres password=dana1234 dbname=todos sslmode=disable"

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return err

}
