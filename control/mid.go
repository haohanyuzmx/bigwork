package control

import (
	"bigwork/foundation"
	"bigwork/jwt"
	"github.com/gin-gonic/gin"
)

func Inspect(c *gin.Context) {
	auth:=c.GetHeader("Authorization")
	if len(auth)<7 {
		c.JSON(200,gin.H{
			"code":500,
			"mess":"请求头错误",
		})
		c.Abort()
		return
	}
	token:=auth[7:]
	var jwt jwt.JWT
	if jwt.Check(token)!=nil{
		foundation.SendMess(c,003,"hacker")
		c.Abort()
		return
	}
	c.Set("username",jwt.Payload.Username)
	//fmt.Println(jwt.Payload)
	c.Next()
}
