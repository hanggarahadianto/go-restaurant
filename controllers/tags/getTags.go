package controllers

import (
	"net/http"
	"restaurant/db"
	"restaurant/models"

	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
	var tagList []models.Tags
	result := db.DB.Debug().Find(&tagList)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   tagList,
	})

}
