package controllers

import (
	"net/http"
	"restaurant/db"
	"restaurant/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	id := c.Param("id")
	var orderData models.Order

	if err := c.ShouldBindJSON(&orderData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()
	newOrder := models.Order{
		Name:          orderData.Name,
		Phone:         orderData.Phone,
		Date:          orderData.Date,
		Time:          orderData.Time,
		Restaurant_ID: id,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	result := db.DB.Debug().Create(&newOrder)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newOrder,
	})

}
