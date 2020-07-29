package main

import (
	"bigwork/control"
	"github.com/gin-gonic/gin"

)

func main()  {
	r:=gin.Default()
	r.POST("login",control.Login)
	r.POST("register",control.Register)
	h:=r.Group("user",control.Inspect)
	h.POST("up",control.Up)
	h.GET("show",control.Show)
	h.POST("share",control.Share)
	r.Static("down","file")
	r.Run()
}
