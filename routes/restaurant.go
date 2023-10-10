package routes

import (
	restaurantController "restaurant/controllers/restaurant"
	"restaurant/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRestaurantRouter(r *gin.Engine) {
	restaurant := r.Group("/restaurant")
	{
		restaurant.GET("/get", restaurantController.GetRestaurants)
		restaurant.POST("/create", middlewares.FileUploadMiddleware(), restaurantController.CreateRestaurant)
		restaurant.GET("getById/:id", restaurantController.RestaurantById)
	}
}
