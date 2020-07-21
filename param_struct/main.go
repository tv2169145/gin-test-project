package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

//type MyTime time.Time
//
//func (mt *MyTime) UnmarshalJSON(bs []byte) error {
//	var s string
//	err := json.Unmarshal(bs, &s)
//	if err != nil {
//		return err
//	}
//	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
//	if err != nil {
//		return err
//	}
//	*mt = MyTime(t)
//	return nil
//}
//
//func (mt MyTime) MarshalJSON() ([]byte, error) {
//	var stamp = fmt.Sprintf("\"%s\"", time.Time(mt).Format("2006-01-02"))
//	return []byte(stamp), nil
//}

//--------------------------------------------------------------------------

//type MyTime struct {
//	time.Time
//}
//
//func (ct *MyTime) UnmarshalJSON(b []byte) error {
//	var err error
//	s := strings.Trim(string(b), "\"")
//	if s == "null" {
//		ct.Time = time.Time{}
//		return err
//	}
//	ct.Time, err = time.Parse("2006-01-02", s)
//	return err
//}
//
//func (ct *MyTime) MarshalJSON() ([]byte, error) {
//	if ct.Time.UnixNano() == nilTime {
//		return []byte("null"), nil
//	}
//	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format("2006-01-02"))), nil
//}
//
//var nilTime = (time.Time{}).UnixNano()
//func (ct *MyTime) IsSet() bool {
//	return ct.UnixNano() != nilTime
//}


type Person struct {
	Name string `json:"name" form:"name"`
	Age int `json:"age" form:"age"`
	Birthday time.Time `json:"birthday" form:"birthday" time_format:"2006-01-02" time_utc:"1" `
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
