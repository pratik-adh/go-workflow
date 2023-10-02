package routes

import (
	"CRUD/api/controller"
	"CRUD/api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controller.RegisterNewUser)

	r.POST("/login", middleware.RateLimiter(), func(ctx *gin.Context) {
		token := controller.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"status": 200,
				"access-token":  token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  401,
				"message": "User not registered.",
			})
		}
	})

	apiRoutes := r.Group("/user-api", middleware.AuthorizeJWT())
	{
		apiRoutes.GET("user", controller.GetUsers)
		apiRoutes.POST("user", controller.CreateUser)
		apiRoutes.GET("user/:id", controller.GetUserByID)
		apiRoutes.PUT("user/:id", controller.UpdateUser)
		apiRoutes.DELETE("user/:id", controller.DeleteUser)
	}
	return r
}
