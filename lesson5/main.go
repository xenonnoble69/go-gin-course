package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// User struct = database model
type User struct {
    ID   uint   `json:"id" gorm:"primaryKey"`
    Name string `json:"name"`
}

func main(){
	r:=gin.Default()
	db,err:=gorm.Open(sqlite.Open("app.db"),&gorm.Config{})
	if err!=nil{
		panic("failed to connect with server")
	}
	db.AutoMigrate(&User{})
	r.POST("/users",func(c *gin.Context){
		var newuser User
		if err:=c.BindJSON(&newuser); err!=nil{
			c.JSON(400,gin.H{
				"invalid":"input",
			})
			return
		}
		db.Create(&newuser)
		c.JSON(200,gin.H{
				"message":"user saved",
				"user":newuser,
			})
	})


	r.GET("/users",func (c *gin.Context)  {
		var user []User
		db.Find(&user)
		c.JSON(200,user)
	})
	r.Run()
}