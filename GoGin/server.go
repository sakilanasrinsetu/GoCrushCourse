package main

import (
	 "github.com/gin-gonic/gin"
	 "gitlab.com/paymenticreview/golang-gin-poc/service"
)

var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(VideoService)
)

func main() {
	server := gin.Default()

	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message" : "OK",
	// 	})
	// })
	server.GET("/posts", func(ctx *gin.Context){
		videoController.FindAll()
	})
	server.Run(":8000")
}