package main

import (
	"Go_shortener/src/models"
	"Go_shortener/src/router"
	"Go_shortener/src/setup"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("unable to load .env")
	}
	r := gin.Default()

	db := setup.GetDB()
	db.AutoMigrate(&models.Goly{})
	db.AutoMigrate(&models.User{})

	users := r.Group("/users")
	golies := r.Group("/golies")
	redirect := r.Group("/r")
	router.RegisterRoutesUser(users)
	router.RegisterRoutesGoly(golies)
	router.RegisterRoutesRedirect(redirect)

	r.Run(":8080")
}
