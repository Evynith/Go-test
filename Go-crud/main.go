package main

import (
	"fmt"
	m "go-crud/models"

	"time"

	u "go-crud/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Init API!")
	router := gin.Default()
	router.GET("/users", u.GetUsers)
	router.POST("/users", u.PostUsers)

	router.Run("localhost:8080")
}

var users = []m.User{
	{
		Name:      "Evy",
		Email:     "evynith@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Name:      "Evy2",
		Email:     "evynith2@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Name:      "Evy3",
		Email:     "evynith3@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
