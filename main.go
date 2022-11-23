package main

import (
	"instagram-service/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/user", controller.CreateUserHandler)
	// r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", controller.UploadPhotoHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
