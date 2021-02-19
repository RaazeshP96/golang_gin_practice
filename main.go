package main

import (
	"os"

	controller "github.com/RaazeshP96/golang_gin_practice/controllers"
	"github.com/RaazeshP96/golang_gin_practice/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	godotenv.Load(".env")
	route := gin.Default()
	route.GET("/videos", func(c *gin.Context) {
		c.JSON(200,
			VideoController.FindAll())
	})
	route.POST("/videos", func(c *gin.Context) {
		c.JSON(200,
			VideoController.Save(c))
	})
	route.Run(":" + os.Getenv("PORT"))
}
