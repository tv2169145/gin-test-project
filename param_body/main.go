package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context){
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		// 若前面已先讀了body, 這裡需要再回寫才可取得form body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))

		message := c.PostForm("message")
		firstname := c.DefaultQuery("first_name", "default")
		lastname := c.DefaultQuery("last_name", "default")
		fmt.Println(message, firstname, lastname)
		c.String(200, string(b))

	})
	r.Run()
}
