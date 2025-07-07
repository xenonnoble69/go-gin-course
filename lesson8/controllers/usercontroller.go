package controllers

import (
	"lesson8/config"
	"lesson8/models"

	"github.com/gin-gonic/gin"
)
func CreateUser(c *gin.Context) {
    var user models.User
	if err:= c.BindJSON(&user) ; err!=nil{
		c.JSON(400,gin.H{
			"invalid":"input",
		})
		return
	}
	config.DB.Create(&user)
	c.JSON(200, gin.H{"message": "User created", "user": user})
}
func GetUsers(c *gin.Context) {
    var user []models.User
	config.DB.Find(&user)
	c.JSON(200, gin.H{"message": "User found", "user": user})

}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	var updated models.User
	if err := c.BindJSON(&updated); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	user.Name = updated.Name
	config.DB.Save(&user)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	config.DB.Delete(&user)
	c.JSON(200, gin.H{"message": "User deleted"})
}