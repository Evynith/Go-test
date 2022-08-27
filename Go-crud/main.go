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

	private1 := router.Group("/")
	private1.Use(authHandler.AuthorizeJWT())
	{
		private1.GET("/users", u.GetUsers)
	}

	var typeUser string = "admin"
	private2 := router.Group("/")
	private2.Use(authHandler.AuthorizeJWT(), authHandler.OnlyAdmin(typeUser))
	{
		private2.GET("/users/:id", u.GetUser)
		private2.POST("/users", u.PostUsers)
		private2.DELETE("/users/:id", u.DeleteUser)
		private2.PUT("/users/:id", u.PutUser)
	}

	router.Run("localhost:8080")

}
