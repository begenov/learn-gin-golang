package api

import (
	"net/http"

	"github.com/begenov/learn-gin-golang/controller"
	"github.com/begenov/learn-gin-golang/dto"
	"github.com/gin-gonic/gin"
)

type VideoApi struct {
	loginController controller.LoginController
	videoController controller.VideoController
}

func NewVideoAPI(loginController controller.LoginController,
	videoController controller.VideoController) *VideoApi {
	return &VideoApi{
		loginController: loginController,
		videoController: videoController,
	}
}

func (api *VideoApi) Authenticate(ctx *gin.Context) {
	token := api.loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, &dto.JWT{
			Token: token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, &dto.Response{
			Message: "Not Authorized",
		})
	}
}

func (api *VideoApi) GetVideos(ctx *gin.Context) {
	ctx.JSON(200, api.videoController.FindAll())
}

func (api *VideoApi) CreateVideo(ctx *gin.Context) {
	err := api.videoController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",
		})
	}
}

func (api *VideoApi) UpdateVideo(ctx *gin.Context) {
	err := api.videoController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",
		})
	}
}

func (api *VideoApi) DeleteVideo(ctx *gin.Context) {
	err := api.videoController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",
		})
	}
}
