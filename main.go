package main

import (
	"Go_shortener/src/models"
	"Go_shortener/src/router"
	"Go_shortener/src/setup"
	"fmt"
	"net/http"

	_ "Go_shortener/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title     Gingo Bookstore API
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("unable to load .env")
	}

	r := gin.Default()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://*", "http://*"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS", "DELETE"},
	// 	AllowHeaders:     []string{"Content-Length", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Origins"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	db := setup.GetDB()
	db.AutoMigrate(&models.Goly{})
	db.AutoMigrate(&models.User{})

	// router for homepage
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title":   "Welcome to Goly Shortener API",
			"Status":  "currently running",
			"message": "To view all public golies head to the /golies"})
	})

	r.Static("/swagger", "./swagger")

	// docs.SwaggerInfo.BasePath = "/api/v1"
	// r.Static("/swagger", "./docs/swagger.json")

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// other router families
	users := r.Group("/users")
	golies := r.Group("/golies")
	redirect := r.Group("/r")

	// register the routes
	router.RegisterRoutesUser(users)
	router.RegisterRoutesGoly(golies)
	router.RegisterRoutesRedirect(redirect)

	r.Run(":8080")
}
