package controllers

import (
	"net/http"
	"restaurant/db"
	"restaurant/models"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	var orderList []models.Order

	result := db.DB.Debug().Find(&orderList)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   orderList,
	})

}
