package setup

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
)

func GetDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("unable to load .env")
	}
	DB_NAME := os.Getenv("DB_NAME")
	DB_INFO := os.Getenv("DB_INFO")
	fmt.Println(DB_INFO, DB_NAME)
	// db, err := gorm.Open("mysql", "root:secret@tcp(mysql:3306)/altsch-go?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(DB_NAME, DB_INFO)
	if err != nil {
		panic(err)
	}
	return db
}

func InitCache() *cache.Cache {
	// Initialize your cache instance
	cache := cache.New(5*time.Minute, 10*time.Minute)

	// Perform any cache initialization tasks here

	return cache
}
