package controller

import (
	"fmt"
	"net/http"

	"CRUD/models"
	"CRUD/service"

	"github.com/gin-gonic/gin"
)

//Register User
func RegisterNewUser(c *gin.Context) {
	var user models.Registration
	c.BindJSON(&user)
	err := service.RegisterNewUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "registered": true, "data": user})
	}
}
