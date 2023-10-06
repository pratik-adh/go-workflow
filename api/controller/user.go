package controller

import (
	"fmt"
	"net/http"

	"CRUD/api/service"
	"CRUD/constants"
	"CRUD/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var user []models.User
	err := service.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func CreateUser(c *gin.Context) {
	var user models.User
	
	tokenString := c.GetHeader("Authorization")
    token, err := jwt.ParseWithClaims(tokenString, &models.AuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(constants.SecretKey), nil
    })
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Check if the token is valid
    if claims, ok := token.Claims.(*models.AuthCustomClaims); ok && token.Valid {
		fmt.Println(claims.ID, claims.Email, constants.SecretKey)
        reportingID := claims.ID
		user.ReportingTo = reportingID
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
    }

	c.BindJSON(&user)

	fmt.Println(user)
	
	errWhileCreatingNewUser := service.CreateUser(c, &user)
	if errWhileCreatingNewUser != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
	}
}

func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := service.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := service.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = service.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := service.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
