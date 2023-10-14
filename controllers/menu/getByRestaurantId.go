package controllers

import (
	"net/http"
	"restaurant/db"
	"restaurant/models"

	"github.com/gin-gonic/gin"
)

func GetMenuByRestaurantId(c *gin.Context) {

	restaurant_id := c.Param("id")

	if restaurant_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'id' parameter"})
		return
	}

	var menu []models.Menu

	result := db.DB.Debug().Find(&menu, "restaurant_id = ?", restaurant_id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "tags doesn't exist",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   menu,
	})
}
