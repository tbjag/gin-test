package main

import (
	"example/web-service-gin/controller"
	"example/web-service-gin/middleware"
	"example/web-service-gin/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
