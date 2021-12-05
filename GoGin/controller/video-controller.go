package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/paymenticreview/golang-gin-poc/entity"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}


func New(service service.VideoService) VideoController {
	return  controller{
		service:service,
	}
}

func  (c *control) findAll() []entity.Video
Save(ctx *gin.Context)