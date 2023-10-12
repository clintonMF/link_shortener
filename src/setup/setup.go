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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASS, DB_HOST, DB_NAME)

	db, err := gorm.Open(DB_TYPE, dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func InitCache() *cache.Cache {
	cache := cache.New(10*time.Hour, 24*time.Hour)

	return cache
}
