package controllers

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

// func CreateUser(c *gin.Context) {
// 	var user models.User
// 	err := c.BindJSON(&user)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusBadRequest)
// 		return
// 	}

// 	models.CreateUser(&user)
// 	c.JSON(http.StatusCreated, user)
// }

// func GetUser(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user models.User
// 	userId, _ := strconv.Atoi(id)
// 	user = models.GetUser(userId)
// 	c.JSON(http.StatusOK, user)
// }

func UploadBangke(c *gin.Context) {
	id := c.Params.ByName("id")

	// filename, ok := c.Get("filePath")
	// if !ok {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "filename not found"})
	// }

	file, ok := c.Get("file")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldnt find file in request"})
		return
	}
	fmt.Println(file)
	fmt.Println(id)

	// // upload file
	// imageUrl, err := utils.UploadToCloudinary(file.(multipart.File), filename.(string))
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// userId, _ := strconv.Atoi(id)
	// update := map[string]string{
	// 	"image_url": imageUrl,
	// }
	// updatedUser := models.UpdateUser(userId, update)
	// c.JSON(http.StatusOK, gin.H{"data": updatedUser})
	// return
}
