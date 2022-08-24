package user_controller

import (
	u "go-crud/services/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	usuarios, _ := u.Read()
	c.IndentedJSON(http.StatusOK, usuarios)
}
