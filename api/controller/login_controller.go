package controller

import (
	"CRUD/api/service"
	"CRUD/models"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) string {
	var credential models.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "No data found"
	}
	isUserAuthenticated := service.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return service.JWTAuthService().GenerateToken(credential.Email, credential.Password)
	}
	return ""
}
