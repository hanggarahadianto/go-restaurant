package routes

import (
	tagController "restaurant/controllers/tags"

	"github.com/gin-gonic/gin"
)

func SetupTagRouter(r *gin.Engine) {
	tag := r.Group("/tag")
	{
		tag.GET("/get", tagController.GetTags)
		tag.POST("/create/:id", tagController.CreateTag)
		tag.GET("/getbyid/:id", tagController.GetTagsById)
		tag.GET("/getbyRestaurantId/:id", tagController.GetTagsByRestaurantId)

	}
}
