package controllers

import (
	"mime/multipart"
	"net/http"
	"restaurant/db"
	"restaurant/models"
	uploadClaudinary "restaurant/utils/cloudinary-folder"

	"github.com/gin-gonic/gin"
)

func CreateMenu(c *gin.Context) {

	id := c.Param("id")

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
	imageUrl, err := uploadClaudinary.UploadtoMenuFolder(file.(multipart.File), filename.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newMenu := models.Menu{
		Title:         c.Request.PostFormValue("title"),
		Content:       c.Request.PostFormValue("content"),
		Image:         imageUrl,
		Restaurant_ID: id,
	}

	result := db.DB.Debug().Create(&newMenu)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newMenu,
	})

}
