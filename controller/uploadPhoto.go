package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadPhotoHandler(context *gin.Context) {
	imageFile, _ := context.FormFile("image")
	userId := context.PostForm("userid")

	if userId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "userid missing "})
		return
	}

	if err := context.SaveUploadedFile(imageFile, "images/"+imageFile.Filename); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to save the file",
		})
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
