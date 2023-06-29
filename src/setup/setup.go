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
	DB_TYPE := os.Getenv("DB_TYPE")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")

	// db, err := gorm.Open("mysql", "root:secret@tcp(mysql:3306)/altsch-go?charset=utf8&parseTime=True&loc=Local")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASS, DB_HOST, DB_NAME)

	db, err := gorm.Open(DB_TYPE, dsn)
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
