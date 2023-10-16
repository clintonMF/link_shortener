package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" gorm:"unique;not null" binding:"required"`
	Password string `json:"password" gorm:"not null" binding:"required"`
	Golies   []Goly `json:"golies" gorm:"foreignKey:UserID"`
}

func (u *User) CreateUser() (*User, error) {
	db.NewRecord(u)
	err := db.Create(&u).Error
	return u, err
}
