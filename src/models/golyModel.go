package models

import (
	"Go_shortener/src/setup"

	"github.com/jinzhu/gorm"
)

type Goly struct {
	gorm.Model
	Redirect string `json:"redirect" binding:"required"`
	Goly     string `json:"goly" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Custom   bool   `json:"custom"`
	UserID   uint   `json:"userId"`
}

type PublicGoly struct {
	Redirect string
	Goly     string
}

var db *gorm.DB = setup.GetDB()

func (g *Goly) CreateGoly() (*Goly, error) {
	db.NewRecord(g)
	err := db.Create(&g).Error
	return g, err
}
