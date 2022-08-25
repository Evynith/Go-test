package login_controller

import (
	"go-crud/credentials"
	service "go-crud/services/login_service"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) string {
	var jwtService service.JWTService = service.JWTAuthService()
	var credential credentials.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := service.LoginUser(credential.Username, credential.Password)
	if isUserAuthenticated {
		return jwtService.GenerateToken(credential.Username, true)

	}
	return ""
}
