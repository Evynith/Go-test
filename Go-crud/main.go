package main

import (
	"net/http"

	controller "go-crud/controllers/login_controller"
	u "go-crud/controllers/user_controller"

	"github.com/gin-gonic/gin"
	//"golang.org/x/oauth2"
)

func main() {
	router := gin.Default()

	router.GET("/users", u.GetUsers)
	router.GET("/users/:id", u.GetUser)
	router.POST("/users", u.PostUsers)
	router.DELETE("/users/:id", u.DeleteUser)
	router.PUT("/users/:id", u.PutUser)
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

	router.Run("localhost:8080")

}
