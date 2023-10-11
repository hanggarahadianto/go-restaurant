package main

import (
	"fmt"
	"log"
	"restaurant/config"
	"restaurant/db"
	"restaurant/routes"
	"restaurant/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	configure, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables ", err)
	}
	gin.SetMode(gin.ReleaseMode)

	db.InitializeDb(&configure)

	r := gin.Default()
	config.LoadEnv()

	routes.SetupUserRouter(r)
	routes.SetupRestaurantRouter(r)
	routes.SetupAuthRouter(r)
	routes.SetupOrderRouter(r)
	routes.SetupTagRouter(r)
	routes.SetupMenuRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my Go image uploader application! 🚀",
		})
	})

	fmt.Println("server 5000")
	log.Fatal(r.Run(":5000"))

	// err := r.Run(":5000")
	// if err != nil {
	// 	panic(err)
	// }
}
