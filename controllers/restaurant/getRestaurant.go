package controllers

import (
	"net/http"
	"restaurant/db"
	"restaurant/models"

	"github.com/gin-gonic/gin"
)

func GetRestaurants(c *gin.Context) {
	var restaurantList []models.Restaurant

	result := db.DB.Debug().Find(&restaurantList)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   restaurantList,
	})

}
