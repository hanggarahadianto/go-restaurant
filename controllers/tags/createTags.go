package controllers

import (
	"net/http"
	"restaurant/db"
	"restaurant/models"

	"github.com/gin-gonic/gin"
)

func CreateTag(c *gin.Context) {
	id := c.Param("id")
	var tagData models.Tags

	if err := c.ShouldBindJSON(&tagData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTag := models.Tags{
		Title:         tagData.Title,
		Restaurant_ID: id,
	}

	result := db.DB.Debug().Create(&newTag)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newTag,
	})

}
