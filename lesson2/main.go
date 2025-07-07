package main

import "github.com/gin-gonic/gin"


// dymanic routing
func main() {

	r := gin.Default()
	r.GET("/user/:id",func(c *gin.Context) {
		id:=c.Param("id")
		c.JSON(200,gin.H{
			"id":id,
		})
	})
	r.GET("/product/:sku",func(c *gin.Context) {
		id:=c.Param("sku")
		c.JSON(200,gin.H{
			"product":id,
		})
	})
	r.Run()
    
}