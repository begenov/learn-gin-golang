package controller

import (
	"github.com/begenov/learn-gin-golang/entity"
	"github.com/begenov/learn-gin-golang/service"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func NewLoginController(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (c *loginController) Login(ctx *gin.Context) string {
	var credentials entity.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := c.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return c.jwtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
