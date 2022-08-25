package main

import (
	"net/http"

	controller "go-crud/controllers/login_controller"
	u "go-crud/controllers/user_controller"
	service "go-crud/services/login_service"

	"github.com/gin-gonic/gin"
	//"golang.org/x/oauth2"
)

func main() {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	router := gin.Default()

	router.GET("/users", u.GetUsers)
	router.GET("/users/:id", u.GetUser)
	router.POST("/users", u.PostUsers)
	router.DELETE("/users/:id", u.DeleteUser)
	router.PUT("/users/:id", u.PutUser)
	router.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	router.Run("localhost:8080")

}
