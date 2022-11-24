package controller

import (
	"instagram-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchPhotoHandler(context *gin.Context) {
	title := context.Param("title") // path parameter
	photos, err := services.SearchPhotoByTitle(title)
	if err != nil {
		context.JSON(http.StatusNoContent, gin.H{
			"images": photos,
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"images": photos,
	})
}
