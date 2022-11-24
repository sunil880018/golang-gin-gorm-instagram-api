package controller

import (
	"fmt"
	"instagram-service/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchPhotoHandler(context *gin.Context) {
	title := context.Param("title") // path parameter

	fmt.Println(title)

	images := []dto.ResponsePhotoDTO{
		{PhotoId: "12331313", PhotoUrl: "http://ssfa23dsfsfdff", Title: "SunSet"},
		{PhotoId: "423423432", PhotoUrl: "http://ssfa23dsfsfdff", Title: "SunSet"},
	}

	context.JSON(http.StatusOK, gin.H{
		"imageList": images,
	})
}
