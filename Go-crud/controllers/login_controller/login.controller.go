package login_controller

import (
	m "go-crud/models"
	service "go-crud/services/login_service"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) string {
	var jwtService service.JWTService = service.JWTAuthService()
	var credential m.Client
	var token string = ""
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := service.LoginUser(credential.Username, credential.Password)
	if isUserAuthenticated {
		token = jwtService.GenerateToken(credential.Username)
	}
	service.SaveToken(credential.Username, token)

	return token
}
