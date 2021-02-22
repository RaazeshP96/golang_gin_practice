package main

import (
	"io"
	"net/http"
	"os"

	controller "github.com/RaazeshP96/golang_gin_practice/controllers"
	"github.com/RaazeshP96/golang_gin_practice/middlewares"
	"github.com/RaazeshP96/golang_gin_practice/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	godotenv.Load(".env")
	setupLogOutput()
	route := gin.New()
	route.Static("/css", "./templates/css")
	route.LoadHTMLGlob("templates/*.html")
	route.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
	apiRoutes := route.Group("/api")
	{
		apiRoutes.GET("/videos", func(c *gin.Context) {
			c.JSON(200,
				videoController.FindAll())
		})
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest,
					gin.H{"error": err.Error()})

			} else {
				ctx.JSON(http.StatusOK,
					gin.H{"message": "Video Input is Valid!!"})

			}

		})
	}
	viewRoute := route.Group("/views")
	{
		viewRoute.GET("/videos", videoController.ShowAll)
	}

	route.Run(":" + os.Getenv("PORT"))
}
