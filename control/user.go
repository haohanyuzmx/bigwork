package control

import (
	"bigwork/foundation"
	"bigwork/jwt"
	"bigwork/module"
	"github.com/gin-gonic/gin"
	"os"
)

type Userinfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context)  {
	var uf Userinfo
	err:=ctx.BindJSON(&uf)
	if foundation.Wrong(err){
		foundation.SendMess(ctx,001,"绑定出错")
		return
	}
	u:=module.User{
		Username: uf.Username,
		Password: uf.Password,
	}
	err=module.Register(u)
	if foundation.Wrong(err) {
		foundation.SendMess(ctx,002,"账号有问题")
		return
	}
	err = os.Mkdir("file/"+uf.Username, os.ModePerm)
	if foundation.Wrong(err) {
		foundation.SendMess(ctx,006,"创建文件失败")
		return
	}
	err = os.Mkdir("share/"+uf.Username, os.ModePerm)
	if foundation.Wrong(err) {
		foundation.SendMess(ctx,006,"创建文件失败")
		return
	}
	foundation.SendMess(ctx,200,"注册成功")
}
func Login(ctx *gin.Context)  {
	var uf Userinfo
	err:=ctx.BindJSON(&uf)
	if foundation.Wrong(err){
		foundation.SendMess(ctx,001,"绑定出错")
		return
	}
	u:=module.User{
		Username: uf.Username,
		Password: uf.Password,
	}
	err=module.Login(u)
	if foundation.Wrong(err) {
		foundation.SendMess(ctx,002,"账号有问题")
		return
	}
	var j jwt.JWT
	j.New(uf.Username)
	foundation.SendMess(ctx,200,j.Token)
}
