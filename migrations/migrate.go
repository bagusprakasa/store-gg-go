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
}
