package main

import (
	"backend-storegg-go/controllers"
	"backend-storegg-go/helpers"
	"backend-storegg-go/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	helpers.LoadEnv()
	helpers.ConnectDb()
}

func main() {
	r := gin.Default()

	// Auth
	r.POST("/register", controllers.Register)
	r.GET("/verified-email/:id", controllers.VerifiedEmail)
	r.POST("/login", controllers.Login)
	r.GET("/check-token", middleware.AuthMiddleware, controllers.CheckToken)

	// Category
	r.GET("/category", middleware.AuthMiddleware, controllers.CategoryIndex)
	r.POST("/category", middleware.AuthMiddleware, controllers.CategoryStore)
	r.GET("/category/:id", middleware.AuthMiddleware, controllers.CategoryShow)
	r.PUT("/category/:id", middleware.AuthMiddleware, controllers.CategoryUpdate)
	r.DELETE("/category/:id", middleware.AuthMiddleware, controllers.CategoryDestroy)

	// Nominal
	r.GET("/nominal", middleware.AuthMiddleware, controllers.NominalIndex)
	r.POST("/nominal", middleware.AuthMiddleware, controllers.NominalStore)
	r.GET("/nominal/:id", middleware.AuthMiddleware, controllers.NominalShow)
	r.PUT("/nominal/:id", middleware.AuthMiddleware, controllers.NominalUpdate)
	r.DELETE("/nominal/:id", middleware.AuthMiddleware, controllers.NominalDestroy)

	r.Run()
}
