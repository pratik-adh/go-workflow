package middleware

import (
	"CRUD/api/service"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// const BEARER_SCHEMA = "Bearer "
		// authHeader := c.GetHeader("Authorization")
		// tokenString := authHeader[len(BEARER_SCHEMA):]
		
		tokenString := c.GetHeader("Authorization")

		token, err := service.JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println("claims", claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": "Please enter a valid token string."})
		}
	}
}
