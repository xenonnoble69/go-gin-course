package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Load DB config
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	// Connect to DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	// Auto-migrate User table
	db.AutoMigrate(&User{})

	r := gin.Default()

	// POST /user
	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
		db.Create(&user)
		c.JSON(200, gin.H{"message": "User created", "user": user})
	})

	// GET /user
	r.GET("/user", func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.JSON(200, users)
	})
r.DELETE("/user/:id", func(c *gin.Context) {
		var user User
		id := c.Param("id")
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "user not found"})
			return
		}
		db.Delete(&user)
		c.JSON(200, gin.H{"message": "user deleted"}) // ✅ added response
	})

	// PUT: Update user
	r.PUT("/user/:id", func(c *gin.Context) {
		var user User
		id := c.Param("id")
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "user not found"})
			return
		}
		var updated User
		if err := c.BindJSON(&updated); err != nil {
			c.JSON(400, gin.H{"error": "invalid input"}) // ✅ better message
			return
		}
		user.Name = updated.Name
		db.Save(&user)
		c.JSON(200, user)
	})
	r.Run()
}
