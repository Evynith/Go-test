package main

import (
	"net/http"

	controller "go-crud/controllers/login_controller"
	u "go-crud/controllers/user_controller"
	authHandler "go-crud/middleware"

	"github.com/gin-gonic/gin"
	//"golang.org/x/oauth2"
)

func main() {
	router := gin.Default()

	router.POST("/login", func(ctx *gin.Context) {
		token := controller.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
	var typeUser string = "admin"
	private1 := router.Group("/users")
	private1.Use(authHandler.AuthorizeJWT())
	{
		private1.GET("/", u.GetUsers)
		private2 := private1.Group("/")
		private2.Use(authHandler.OnlyAdmin(typeUser))
		{
			private2.GET("/:id", u.GetUser)
			private2.POST("/", u.PostUsers)
			private2.DELETE("/:id", u.DeleteUser)
			private2.PUT("/:id", u.PutUser)
		}
	}

	router.Run("localhost:8080")

}
