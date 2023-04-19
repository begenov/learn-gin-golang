package main

import (
	"io"
	"net/http"
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
	loginController controller.LoginController = controller.NewLoginController(service.NewLoginService(), service.NewJWTService())
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

	mux.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := mux.Group("/api", middlewares.AuthorizeJWT())
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
				ctx.JSON(http.StatusOK, gin.H{"MESSAGE": "VIDEO INPUT IS VALID!!!"})
			}

		})
	}

	viewRoutes := mux.Group("/view")

	{
		viewRoutes.GET("/videos", VideoController.ShowAll)
	}

	mux.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump(),
	)

	mux.Run(":8090")

}
