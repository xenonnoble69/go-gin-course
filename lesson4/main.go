package main

import "github.com/gin-gonic/gin"

type person struct{
	ID int `json:"id"`
	Name string `json:"name"`
}
type product struct{
	Title string `json:"title"`
	Price int `json:"price"`
}

func main() {
	r := gin.Default()
	r.POST("/user",func(c *gin.Context) {
		var newuser person
		if err:=c.BindJSON(&newuser); err!=nil{
			c.JSON(400,gin.H{"error":"invalid input"})
			return
		}
		c.JSON(200,gin.H{"recived_name":newuser.Name,
		"recived_id":newuser.ID,
	})
	})
	r.POST("/product",func(c *gin.Context) {
		var newproduct product
		if err:=c.BindJSON(&newproduct); err!=nil{
			c.JSON(400,gin.H{"error":"invalid input"})
			return
		}
		c.JSON(200,gin.H{"recived_title":newproduct.Title,
		"recived_price":newproduct.Price,
	})
	})
	r.Run()

}