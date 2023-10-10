package routes

import (
	"restaurant/controllers"
	"restaurant/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(r *gin.Engine) {
	user := r.Group("/user")
	{
		// user.POST("/", controllers.CreateUser)
		// user.GET("/:id", controllers.GetUser)
		user.POST("/:id/uploadImage", middlewares.FileUploadMiddleware(), controllers.UploadBangke)
		// user.POST("/:id/uploadImage", controllers.UploadImage)
	}
}
