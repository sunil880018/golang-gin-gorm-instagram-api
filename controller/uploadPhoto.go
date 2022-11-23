package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadPhotoHandler(context *gin.Context) {
	// Multipart form

	imageFile, _ := context.FormFile("image")
	// Upload the file to specific dst.
	if err := context.SaveUploadedFile(imageFile, "images/"+imageFile.Filename); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"image": imageFile.Filename})
}
