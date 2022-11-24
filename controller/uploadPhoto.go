package controller

import (
	"instagram-service/dto"
	"instagram-service/enums"
	"instagram-service/helper"
	"instagram-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadPhotoHandler(context *gin.Context) {
	imageFile, _ := context.FormFile("image")
	var photo dto.PhotoDTO
	photo.Title = imageFile.Filename

	userId := context.PostForm("userid")
	if userId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": enums.BAD_REQUEST})
		return
	}

	services.CreateImage(photo)
	// upload photo on aws s3
	helper.UploadOnS3Bucket(photo)

	services.CreateUserPhotoDetails(photo, userId)
	// upload in file directory
	if err := context.SaveUploadedFile(imageFile, "images/"+imageFile.Filename); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": enums.INTERNAL_SERVER_ERROR})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"image": imageFile.Filename, "userid": userId})
}

// curl --location --request POST 'localhost:8080/upload' \
// --form 'image=@"Your Path/anu.jpg"' \
// --form 'userid="34724y6274"'

// ---------------- upload file with some data in Postman ---------------------------

//  1. POST - localhost:8000
//  2. select Body > form data >
//     i) file   :<path>
//     ii) userid : "userid"

// ---------------------output ----------------------
// {
//     "image": "anu.jpg",
//     "userid": "34724y6274"
// }
