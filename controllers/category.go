package controllers

import (
	"backend-storegg-go/helpers"
	"backend-storegg-go/models"

	"github.com/gin-gonic/gin"
)

func CategoryIndex(c *gin.Context) {
	var model []models.Category
	helpers.DB.Find(&model)

	// Return Success
	c.JSON(200, gin.H{
		"data":    model,
		"message": "Success",
	})
}

func CategoryStore(c *gin.Context) {
	// Data from Req Body
	var request struct {
		Name string
	}
	c.BindJSON(&request)
	model := models.Category{Name: request.Name}
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

func CategoryShow(c *gin.Context) {
	// Get Param
	id := c.Param("id")

	// Find by id
	var model models.Category
	helpers.DB.First(&model, id)

	// Return Success
	c.JSON(200, gin.H{
		"data":    model,
		"message": "Success",
	})
}

func CategoryUpdate(c *gin.Context) {
	// Get Param
	id := c.Param("id")

	// Data from Req Body
	var request struct {
		Name string
	}
	c.BindJSON(&request)

	// Find data were update
	var model models.Category
	helpers.DB.First(&model, id)

	// Update
	helpers.DB.Model(&model).Updates(models.Category{
		Name: request.Name,
	})

	// Return Success
	c.JSON(200, gin.H{
		"message": "Success",
		"data":    model,
	})
}

func CategoryDestroy(c *gin.Context) {
	// Get Param
	id := c.Param("id")

	// Find by id
	var model models.Category
	helpers.DB.Delete(&model, id)

	// Return Success
	c.JSON(200, gin.H{
		// "data":    model,
		"message": "Success",
	})
}
