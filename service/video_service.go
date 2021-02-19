package service

import "github.com/RaazeshP96/golang_gin_practice/models"

type VideoService interface {
	Save(models.Video) models.Video

	FindAll() []models.Video
}

type videoService struct {
	videos []models.Video
}

// Create a new instance for the video

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video models.Video) models.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []models.Video {
	return service.videos

}
