package main

import (
	"backend-storegg-go/helpers"
	"backend-storegg-go/models"
)

func init() {
	helpers.LoadEnv()
	helpers.ConnectDb()
}

func main() {
	helpers.DB.AutoMigrate(&models.Category{})
	helpers.DB.AutoMigrate(&models.Nominal{})
	helpers.DB.AutoMigrate(&models.User{})
}
