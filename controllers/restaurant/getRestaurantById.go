package controllers

import (
	"net/http"
	"restaurant/db"
	"restaurant/models"

	"github.com/gin-gonic/gin"
)

func RestaurantById(c *gin.Context) {
	restaurantId := c.Param("id")

	var restaurant models.Restaurant

	result := db.DB.Debug().First(&restaurant, "id = ?", restaurantId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "restaurant id doesn't exist",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   restaurant,
	})
}
