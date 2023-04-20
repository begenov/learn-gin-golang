package main

import (
	"io"
	"net/http"
	"os"

	"github.com/begenov/learn-gin-golang/controller"
	"github.com/begenov/learn-gin-golang/middlewares"
	"github.com/begenov/learn-gin-golang/repository"
	"github.com/begenov/learn-gin-golang/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoRepositroy repository.VedeoRepository = repository.NewVideoRepository()
	videoService    service.VedioService       = service.New(videoRepositroy)
	VideoController controller.VideoController = controller.New(videoService)
)

func sendLogOutput() {
	f, _ := os.Create("gin.log.txt")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	sendLogOutput()

	mux := gin.New()

	mux.Static("/css", "./templates/css")

	mux.LoadHTMLGlob("templates/*.html")

	apiRoutes := mux.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, VideoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := VideoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"ERROR": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"MESSAGE": "vIDEO INPUT IS VALID!!!"})
			}
		})
		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := VideoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"ERROR": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"MESSAGE": "vIDEO INPUT IS VALID!!!"})
			}
		})
		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := VideoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"ERROR": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"MESSAGE": "vIDEO INPUT IS VALID!!!"})
			}
		})
	}

	viewRoutes := mux.Group("/view")

	mux.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump(),
	)
	{
		viewRoutes.GET("/videos", VideoController.ShowAll)
	}

	mux.Run(":8090")
}
