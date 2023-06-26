package models

import (
	"Go_shortener/src/setup"
	"Go_shortener/src/utils"

	"github.com/jinzhu/gorm"
)

type Goly struct {
	gorm.Model
	Redirect string `json:"redirect"`
	Goly     string `json:"goly" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Custom   bool   `json:"custom"`
	Public   bool   `json:"public" gorm:"default:false"`
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

func GetGolies() ([]Goly, error) {
	var golies []Goly
	result := db.Find(&golies)

	if result.Error != nil {
		return nil, result.Error
	}

	return golies, nil
}

func GetGolyByID(ID uint) (*Goly, error) {
	goly := &Goly{}
	if err := db.First(goly, ID).Error; err != nil {
		return nil, utils.ErrNotFound("Goly", int(ID))
	}
	return goly, nil
}

func GetGolyByURL(url string) (*Goly, error) {
	goly := &Goly{}
	if err := db.Where("Goly = ?", url).First(goly).Error; err != nil {
		return nil, utils.ErrNotFoundUrl("Goly", url)
	}
	return goly, nil
}

func GetGoliesByUserID(userID uint) ([]Goly, error) {
	var golies []Goly
	if err := db.Where("user_id = ?", userID).Find(&golies).Error; err != nil {
		return nil, err
	}

	return golies, nil
}

func GetPublicGolies() ([]PublicGoly, error) {
	var publicGolies []PublicGoly
	var golies []Goly
	if err := db.Where("public = ?", true).Find(&golies).Error; err != nil {
		return nil, err
	}
	for _, goly := range golies {
		pubGo := PublicGoly{
			Redirect: goly.Redirect,
			Goly:     goly.Goly,
		}

		publicGolies = append(publicGolies, pubGo)
	}
	return publicGolies, nil
}

func UpdateGoly(goly *Goly) error {
	if err := db.Save(&goly); err != nil {
		return err.Error
	}

	return nil
}

func DeleteGoly(golyID uint) error {
	var goly Goly
	err := db.Find(&goly, golyID).Error

	if err != nil {
		return utils.ErrNotFound("course", int(golyID))
	}

	if err := db.Where("id = ?", golyID).Delete(&Goly{}).Error; err != nil {
		return err
	}
	return nil
}
