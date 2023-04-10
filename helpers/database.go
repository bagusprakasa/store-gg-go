package helpers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var DB *gorm.DB
var DB *gorm.DB

func ConnectDb() {
	var err error
	dsn := os.Getenv("URL_DB")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}
