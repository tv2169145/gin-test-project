package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context){
		firstname := c.Query("first_name")
		lastname := c.DefaultQuery("last_name", "last_default_name")
		c.JSON(200, gin.H{
			"firstname":firstname,
			"lastname":lastname,
		})
	})
	r.Run()
}
