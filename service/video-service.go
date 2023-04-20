package service

import (
	"github.com/begenov/learn-gin-golang/entity"
	"github.com/begenov/learn-gin-golang/repository"
)

type VedioService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
	Update(entity.Video)
	Delete(entity.Video)
}

type videoService struct {
	videoRepository repository.VedeoRepository
}

func NewService(repo repository.VedeoRepository) VedioService {
	return &videoService{
		videoRepository: repo,
	}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videoRepository.Save(video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()
}

func (service *videoService) Update(video entity.Video) {
	service.videoRepository.Update(video)
}

func (service *videoService) Delete(video entity.Video) {
	service.videoRepository.Delete(video)
}
