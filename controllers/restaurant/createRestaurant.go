package controllers

import (
	"mime/multipart"
	"net/http"
	"restaurant/db"
	"restaurant/models"
	"restaurant/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(c *gin.Context) {

	filename, ok := c.Get("filePath")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "filename not found",
			"data":  filename,
		})
	}
	file, ok := c.Get("file")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
		})
	}

	// upload file
	imageUrl, err := utils.UploadToCloudinary(file.(multipart.File), filename.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	newRestaurant := models.Restaurant{
		Name:      c.Request.PostFormValue("name"),
		Content:   c.Request.PostFormValue("content"),
		Phone:     c.Request.PostFormValue("phone"),
		CreatedAt: now,
		UpdatedAt: now,
	}

	newRestaurant.Image = imageUrl

	result := db.DB.Debug().Create(&newRestaurant)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newRestaurant,
	})

}
