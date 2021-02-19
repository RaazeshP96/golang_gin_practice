package controllers

import (
	"github.com/RaazeshP96/golang_gin_practice/models"
	"github.com/RaazeshP96/golang_gin_practice/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Findall() []models.Video
	Save(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller{
		service: service,
	}
}
func (c *controller) Save(ctx *gin.Context) {
	var video models.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
func (c *controller) FindAll() models.Video {

	return c.service.FindAll()
}
