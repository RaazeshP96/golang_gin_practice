package main

import (
	"os"

	"github.com/RaazeshP96/golang_gin_practice/controllers"
	"github.com/RaazeshP96/golang_gin_practice/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	videoService    service.VideoService        = service.New()
	VideoController controllers.VideoController = controllers.New()
)

func main() {
	godotenv.Load(".env")
	route := gin.Default()
	route.GET("/videos", func(c *gin.Context) {
		c.JSON(200,
			VideoController.Findall())
	})
	route.POST("/videos", func(c *gin.Context) {
		c.JSON(200,
			videoController.Save(c))
	})
	route.Run(":" + os.Getenv("PORT"))
}
