package main

import (
	"instagram-service/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20

	r.POST("/user", controller.CreateUserHandler)
	r.POST("/upload", controller.UploadPhotoHandler)
	r.GET("/image/:title", controller.SearchPhotoHandler)
	r.POST("/follow", controller.UserFollowingHandler)
	r.Run()
}
