package routes

import (
	authController "restaurant/controllers/auth"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", authController.Login)

	}
}
