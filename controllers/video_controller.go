package controller

import (
	"github.com/RaazeshP96/golang_gin_practice/models"
	"github.com/RaazeshP96/golang_gin_practice/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []models.Video
	Save(ctx *gin.Context) models.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []models.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) models.Video {
	var video models.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
