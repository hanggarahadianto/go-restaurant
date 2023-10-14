package routes

import (
	menuController "restaurant/controllers/menu"
	"restaurant/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupMenuRouter(r *gin.Engine) {
	menu := r.Group("/menu")
	{
		menu.GET("/get", menuController.GetMenus)
		menu.POST("/create/:id", middlewares.FileUploadMiddleware(), menuController.CreateMenu)
		menu.GET("/getbyRestaurantId/:id", menuController.GetMenuByRestaurantId)

	}
}
