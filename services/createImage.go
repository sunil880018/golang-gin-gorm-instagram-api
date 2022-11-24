package services

import (
	"fmt"
	"instagram-service/database"
	"instagram-service/dto"
	"instagram-service/models"
)

func CreateImage(photo dto.PhotoDTO) {
	var PhotoObj models.Photos

	PhotoObj.PhotoUrl = photo.PhotoUrl
	PhotoObj.Title = photo.Title
	photoRecord := database.Database.Create(&PhotoObj)

	if photoRecord != nil {
		panic(photoRecord.Error)
	}
	fmt.Println(PhotoObj)
}
