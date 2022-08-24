package user_controller

import (
	m "go-crud/models"
	u "go-crud/services/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	usuarios, err := u.Read()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, usuarios)
}

func PostUsers(c *gin.Context) {
	var newUser m.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	err := u.Create(newUser)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newUser)
}
