package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
"gorm.io/driver/sqlite"
)

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic("db not created")
	}

	db.AutoMigrate(&User{})

	// POST: Create user
	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"invalid": "input", // ✅ typo fixed
			})
			return
		}
		db.Create(&user)
		c.JSON(200, gin.H{
			"message": "user created",
			"data":    user,
		})
	})

	// GET: All users
	r.GET("/user", func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.JSON(200, users)
	})

	// GET: Single user by ID
	r.GET("/user/:id", func(c *gin.Context) {
		var user User
		id := c.Param("id")
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "user not found"})
			return
		}
		c.JSON(200, user)
	})

	// DELETE: Delete user
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
