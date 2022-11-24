package services

import (
	"instagram-service/database"
	"instagram-service/dto"
	"instagram-service/models"
)

func CreateImage(photo dto.PhotoDTO) (models.Photos, error) {
	var PhotoObj models.Photos

	PhotoObj.PhotoUrl = photo.PhotoUrl
	PhotoObj.Title = photo.Title
	err := database.Database.Create(&PhotoObj).Error

	if err != nil {
		return models.Photos{}, err
	}
	return PhotoObj, nil
}
