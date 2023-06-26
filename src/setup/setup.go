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
	db_name := os.Getenv("db_name")
	DB_INFO := os.Getenv("DB_INFO")
	db, err := gorm.Open(db_name, DB_INFO)
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
