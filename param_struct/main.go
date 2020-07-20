package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Person struct {
	Name string `json:"name" form:"name"`
	Age int `json:"age" form:"age"`
	Birthday time.Time `json:"birthday" form:"birthday" time_format:"2006-01-02" `
}

func main() {
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)
	r.Run()
}

func testing(c *gin.Context) {
	var p Person
	if err := c.ShouldBind(&p); err == nil {
		c.String(200, "%v", p)
	} else {
		c.String(400, "person bind error: %v", err)
	}
}
