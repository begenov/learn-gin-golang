package main

import (
	"io"
	"os"

	"github.com/begenov/learn-gin-golang/controller"
	"github.com/begenov/learn-gin-golang/middlewares"
	"github.com/begenov/learn-gin-golang/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VedioService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func sendLogOutput() {
	f, _ := os.Create("gin.log.txt")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func main() {

	sendLogOutput()

	mux := gin.New()

	mux.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	mux.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())

	})

	mux.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))

	})

	mux.Run(":8080")

}
