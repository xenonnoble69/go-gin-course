package main

import "github.com/gin-gonic/gin"
// unnderstanding links and query param
func main() {
	r := gin.Default()
	r.GET("/search",func(c *gin.Context) {
		name:=c.Query("name")
		id:=c.Query("id")
        c.JSON(200,gin.H{
			"id":id,
			"name":name,
		})
	})
 r.Run()
}