package routes

import (
	orderController "restaurant/controllers/order"

	"github.com/gin-gonic/gin"
)

func SetupOrderRouter(r *gin.Engine) {
	auth := r.Group("/order")
	{
		auth.GET("/get", orderController.GetOrders)
		auth.POST("/create/:id", orderController.CreateOrder)

	}
}
