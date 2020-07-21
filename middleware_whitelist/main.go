package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(whitelist())
	r.GET("/test", func(c *gin.Context){
		c.String(200, "hello")
	})
	r.Run()
}

func whitelist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.1",
			"::1",
		}
		flag := false
		clientIp := c.ClientIP()
		for _, ip := range ipList {
			if clientIp == ip {
				flag = true
				break
			}
		}
		if !flag {
			c.String(401, "ip not allow")
			return
		}
	}
}
