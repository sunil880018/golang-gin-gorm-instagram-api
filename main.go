package main

import (
	"instagram-service/controller"
	"instagram-service/database"
	"instagram-service/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.Photos{}) // automatic tables created in postgres db
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.UserFollower{})
	database.Database.AutoMigrate(&models.UserPhotoDetails{})
}

func serveApplication() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20

	r.POST("/user", controller.CreateUserHandler)
	r.POST("/upload", controller.UploadPhotoHandler)
	r.GET("/image/:title", controller.SearchPhotoHandler)
	r.POST("/follow", controller.UserFollowingHandler)
	r.Run()
}

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}
