package services

import (
	"instagram-service/database"
	"instagram-service/models"
)

func SearchPhotoByTitle(title string) ([]models.Photos, error) {
	var photos []models.Photos
	err := database.Database.Find(&photos, title).Error
	if err != nil {
		return []models.Photos{}, err
	}
	return photos, nil
}
