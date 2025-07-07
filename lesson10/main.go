// main.go
package main

import (
	"lesson10/config"
	"lesson10/routes"
    //   "github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect() // Connects to DB

	r := gin.Default()

	routes.AuthRoutes(r) // Register /auth routes

	r.Run() // :8080
}
