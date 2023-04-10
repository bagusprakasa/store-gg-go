package controllers

import (
	"backend-storegg-go/helpers"
	"backend-storegg-go/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// Data from Req Body
	var request struct {
		Name     string
		Email    string
		Password string
	}

	if c.BindJSON(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to hash password",
			"message": err.Error(),
		})
		return
	}

	// Insert User
	model := &models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hash),
	}

	result := helpers.DB.Create(&model)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to insert",
			"message": result.Error,
		})
		return
	}
	// Return Success
	c.JSON(200, gin.H{
		"data":    time.Now(),
		"message": "Success",
	})
}

func VerifiedEmail(c *gin.Context) {
	// Get Param
	id := c.Param("id")

	// Find data were update
	var model models.User
	helpers.DB.First(&model, id)

	// Update
	helpers.DB.Model(&model).Updates(models.User{
		EmailVerifiedAt: time.Now(),
	})

	// Return Success
	c.JSON(200, gin.H{
		"data":    model,
		"message": "Success",
	})
}

func Login(c *gin.Context) {
	// Data from Req Body
	var request struct {
		Email    string
		Password string
	}

	if c.BindJSON(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
	}

	var model models.User
	helpers.DB.First(&model, "email = ?", request.Email)

	if model.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to get user from database",
			"message": "User not found",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Login failed",
			"message": "Invalid password or email",
		})

		return
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": model.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Get token after login
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Login failed",
			"message": "Failed create token",
		})

		return
	}

	// Return Success
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*0.5, "", "", false, true)
	c.JSON(200, gin.H{
		"message": "Success",
	})
}

func CheckToken(c *gin.Context) {
	c.JSON(200, gin.H{
		"token": "Success",
	})
}
