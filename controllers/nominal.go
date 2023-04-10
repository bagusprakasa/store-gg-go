package controllers

import (
	"backend-storegg-go/helpers"
	"backend-storegg-go/models"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func NominalIndex(c *gin.Context) {
	var model []models.Nominal
	helpers.DB.Find(&model)

	// Return Success
	c.JSON(200, gin.H{
		"data":    model,
		"message": "Success",
	})
}

func NominalStore(c *gin.Context) {
	// Data from Req Body
	var request struct {
		Name     string
		Quantity decimal.Decimal
		Price    decimal.Decimal
	}
	c.BindJSON(&request)
	model := models.Nominal{Name: request.Name, Quantity: request.Quantity, Price: request.Price}
	result := helpers.DB.Create(&model)
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return Success
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    model,
	})
}

func NominalShow(c *gin.Context) {
	// Get Param
	id := c.Param("id")

	// Find by id
	var model models.Nominal
	helpers.DB.First(&model, id)

	// Return Success
	c.JSON(200, gin.H{
		"data":    model,
		"message": "Success",
	})
}

func NominalUpdate(c *gin.Context) {
	// Get Param
	id := c.Param("id")

	// Data from Req Body
	var request struct {
		Name     string
		Quantity decimal.Decimal
		Price    decimal.Decimal
	}
	c.BindJSON(&request)

	// Find data were update
	var model models.Nominal
	helpers.DB.First(&model, id)

	// Update
	helpers.DB.Model(&model).Updates(models.Nominal{
		Name:     request.Name,
		Quantity: request.Quantity,
		Price:    request.Price,
	})

	// Return Success
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    model,
	})
}

func NominalDestroy(c *gin.Context) {
	// Get Param
	id := c.Param("id")

	// Find by id
	var model models.Nominal
	helpers.DB.Delete(&model, id)

	// Return Success
	c.JSON(200, gin.H{
		// "data":    model,
		"message": "Success",
	})
}
