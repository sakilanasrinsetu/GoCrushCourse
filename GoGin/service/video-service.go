package service

import "gitlab.com/pragmaticreviews/golang-gin-poc/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}
type videoService struct {
	videos []entity.Video 
}

func New() VideoService {
	return &videoService{}
}
func (service *videoService) Save(entity.Video) entity.Video{
	service.videos = append(service.Video, video)
	return service.video
}
func (service *videoService) FindAll() []entity.Video{
	return service.video
}