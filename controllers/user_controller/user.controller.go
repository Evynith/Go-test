package user_controller

import (
	m "go-crud/models"
	u "go-crud/services/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	name, _ := c.GetQuery("name")   //contiene string llegada
	date1, _ := c.GetQuery("desde") //fecha(created_at) - ?null
	date2, _ := c.GetQuery("hasta") //fecha(created_at) - ?null

	usuarios, err := u.Read(name, date1, date2)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, usuarios)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := u.ReadOne(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
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

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := u.Delete(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(200)
}

func PutUser(c *gin.Context) {
	id := c.Param("id")
	var newUser m.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	err := u.Update(newUser, id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, newUser)
}
