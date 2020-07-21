package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"time"
)

type Booking struct {
	CheckIn time.Time `form:"check_in" time_format:"2006-01-02" binding:"required,bookabledate"`
	CheckOut time.Time `form:"check_out" time_format:"2006-01-02" binding:"required,gtfield=CheckIn"`
}

func bookableDate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if date, ok := field.Interface().(time.Time); ok {
		now := time.Now()
		if now.Unix() > date.Unix() {
			return false
		}
	}
	return true
}

func main() {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	r.GET("/test", func(c *gin.Context) {
		var booking Booking
		if err := c.ShouldBind(&booking); err == nil {
			c.String(200, "%v", booking)
		} else {
			c.String(400, "valid error: %v", err)
		}
	})
	r.Run()
}
