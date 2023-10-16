package services

import (
	"Go_shortener/src/models"
	"Go_shortener/src/utils"
)

func GetUserByID(ID uint) (*models.User, error) {
	user := &models.User{}
	if err := db.First(user, ID).Error; err != nil {
		return nil, utils.ErrNotFound("user", int(ID))
	}
	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		return nil, utils.ErrNotFoundUser("User", email)
	}
	return user, nil
}
