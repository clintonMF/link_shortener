package models

import (
	"Go_shortener/src/utils"

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

func GetUserByID(ID uint) (*User, error) {
	user := &User{}
	if err := db.First(user, ID).Error; err != nil {
		return nil, utils.ErrNotFound("user", int(ID))
	}
	return user, nil
}

func GetUserByEmail(email string) (*User, error) {
	user := &User{}
	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		return nil, utils.ErrNotFoundUser("User", email)
	}
	return user, nil
}

// func UpdateGoly(goly *Goly) error {
// 	if err := db.Save(&goly); err != nil {
// 		return err.Error
// 	}

// 	return nil
// }

// func DeleteGoly(golyID uint64) error {
// 	var goly Goly
// 	err := db.Find(&goly, golyID).Error

// 	if err != nil {
// 		return utils.ErrNotFound("course", int(golyID))
// 	}

// 	if err := db.Where("id = ?", golyID).Delete(&Goly{}).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
