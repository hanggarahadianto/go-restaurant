package controllers

import (
	"net/http"
	"restaurant/db"
	"restaurant/models"

	"github.com/gin-gonic/gin"
)

func GetTagsByRestaurantId(c *gin.Context) {
	tagId := c.Param("id")
	var tag models.Tags

	result := db.DB.Debug().First(&tag, "id = ?", tagId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "restaurant id doesn't exist",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   tag,
	})
}
