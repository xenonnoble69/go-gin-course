package main

import (
	"lesson8/config"
	// "lesson8/models"
	"lesson8/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.Connect()
	// config.DB.AutoMigrate(&models.User{})

	routes.UserRoutes(r)

	r.Run()
}
