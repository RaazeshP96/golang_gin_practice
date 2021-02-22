package main

import (
	"io"
	"os"

	controller "github.com/RaazeshP96/golang_gin_practice/controllers"
	"github.com/RaazeshP96/golang_gin_practice/middlewares"
	"github.com/RaazeshP96/golang_gin_practice/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	godotenv.Load(".env")
	setupLogOutput()
	route := gin.New()
	route.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())
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
