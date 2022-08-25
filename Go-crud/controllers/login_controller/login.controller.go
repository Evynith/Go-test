package login_controller

import (
	"go-crud/credentials"
	service "go-crud/services/login_service"

	"github.com/gin-gonic/gin"
)

// login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential credentials.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Username, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.Username, true)

	}
	return ""
}
