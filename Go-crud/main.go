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

	private := router.Group("/")
	private.Use(authHandler.AuthorizeJWT())
	{
		private.GET("/users", u.GetUsers)
		private.GET("/users/:id", u.GetUser)
		private.POST("/users", u.PostUsers)
		private.DELETE("/users/:id", u.DeleteUser)
		private.PUT("/users/:id", u.PutUser)
	}

	router.Run("localhost:8080")

}
