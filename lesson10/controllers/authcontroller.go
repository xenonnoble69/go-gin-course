package controllers

import (
	"time"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"lesson10/config"
	"lesson10/models"
	"golang.org/x/crypto/bcrypt"

)
 
func Register(c *gin.Context)  {
// 	var user models.User
// 	if err := c.BindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}
//    var existing models.User
// 	if err := config.DB.Where("email = ?", user.Email).First(&existing).Error; err == nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
// 		return
// 	}
var user models.User
if err:=c.BindJSON(&user); err!=nil{
	c.JSON(http.StatusBadRequest,gin.H{
		"error":"inavlid request",
	})
	return
}
var existinguser models.User
if err:=config.DB.Where("email=?",user.Email).First(&existinguser).Error ; err==nil{
	c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
}
hashpassword,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)// Save user to DB
if err!=nil{
	c.JSON(http.StatusBadRequest,gin.H{
		"error":"inavlid request",
	})
	return
}
user.Password=string(hashpassword)
	config.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User registered", "user": user})
	
}


func Login(c *gin.Context) {
	var requestBody models.User
	var dbUser models.User

	// Bind the JSON
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Find user in DB using email
	if err := config.DB.Where("email = ?", requestBody.Email).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not registered"})
		return
	}

	// Check password (here plain-text match for now, later we’ll hash)
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(requestBody.Password)); err != nil {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
	return
}


	// Generate JWT
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"user_id": dbUser.Id,
	// 	"exp":     time.Now().Add(time.Hour * 24).Unix(), // token expires in 1 day
	// })

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id":dbUser.Id,
		"exp":time.Now().Add(time.Hour*24).Unix(),
	})

	// Secret key from env
	secret := os.Getenv("JWT_SECRET")

	// Sign the token
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	// Send token in response
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
