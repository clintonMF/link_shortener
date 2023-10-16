package services

import (
	"Go_shortener/src/models"
	"Go_shortener/src/setup"
	"Go_shortener/src/utils"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB = setup.GetDB()

func GetGolies() ([]models.Goly, error) {
	var golies []models.Goly
	result := db.Find(&golies)

	if result.Error != nil {
		return nil, result.Error
	}

	return golies, nil
}

func GetGolyByID(ID uint) (*models.Goly, error) {
	goly := &models.Goly{}
	if err := db.First(goly, ID).Error; err != nil {
		return nil, utils.ErrNotFound("Goly", int(ID))
	}
	return goly, nil
}

func GetGolyByURL(url string) (*models.Goly, error) {
	goly := &models.Goly{}
	if err := db.Where("Goly = ?", url).First(goly).Error; err != nil {
		return nil, utils.ErrNotFoundUrl("Goly", url)
	}
	return goly, nil
}

func GetGoliesByUserID(userID uint) ([]models.Goly, error) {
	var golies []models.Goly
	if err := db.Where("user_id = ?", userID).Find(&golies).Error; err != nil {
		return nil, err
	}

	return golies, nil
}

func UpdateGoly(goly *models.Goly) error {
	if err := db.Save(&goly); err != nil {
		return err.Error
	}

	return nil
}

func DeleteGoly(golyID uint) error {
	var goly models.Goly
	err := db.Find(&goly, golyID).Error

	if err != nil {
		return utils.ErrNotFound("course", int(golyID))
	}

	if err := db.Where("id = ?", golyID).Delete(&models.Goly{}).Error; err != nil {
		return err
	}
	return nil
}
