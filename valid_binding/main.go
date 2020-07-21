package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name" form:"name"`
	Age int `json:"age" form:"age" binding:"required,gt=18"`
	Address string `json:"address" form:"address" binding:"required"`
}

func main() {
	r := gin.Default()
	r.Any("/test", func(c *gin.Context) {
		var p Person
		if err := c.ShouldBind(&p); err == nil {
			c.String(200, "%v", p)
		} else {
			c.String(400, "person valid error: %v", err)
		}

	})
	r.Run()
}

