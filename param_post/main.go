package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context){
		name := c.PostForm("name")
		age := c.DefaultPostForm("age", "18")
		c.String(200, fmt.Sprintf("%s / %s\n", name, age))
	})
	r.Run()
}
