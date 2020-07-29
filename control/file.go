package control

import (
	"bigwork/foundation"
	"bigwork/module"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"strings"
)

func Show(ctx *gin.Context)  {
	n,_:=ctx.Get("username")
	m:=make(map[int][]string)
	f:=module.ShowMyFile(n.(string))
	for i, i2 := range f {
		ss:=make([]string,2)
		s0:="真正的："+i2.Uri
		s1:="加密后："+base64.StdEncoding.EncodeToString([]byte(i2.Uri))
		ss[0]=s0
		ss[1]=s1
		m[i+1]=ss
	}
	ctx.JSON(200,m)
}
func Up(ctx *gin.Context)  {
	n,_:=ctx.Get("username")
	fil,err:=ctx.FormFile("file")
	if foundation.Wrong(err) {
		foundation.SendMess(ctx,001,"绑定出错")
		return
	}
	if fil.Size<1024*1024*100 {
		ctx.SaveUploadedFile(fil,"file/"+n.(string)+"/"+fil.Filename)
	}else {
		foundation.Savefile(fil,"file/"+n.(string)+"/"+fil.Filename)
	}
	ff:=module.MyFile{
		Username: n.(string),
		Uri: "localhost:8080/down/"+n.(string)+"/"+fil.Filename,
	}
	err=module.InsertFile(&ff)
	if foundation.Wrong(err) {
		foundation.SendMess(ctx,004,"数据库抽风")
	}
	foundation.SendMess(ctx,200,"成功")
}
func Share(ctx *gin.Context)  {
	n,_:=ctx.Get("username")
	url:=ctx.PostForm("url")
	paths:=strings.Split(url,"/")
	na:=paths[len(paths)-1]
	thing:=strings.Split(na,".")
	qc:="share/"+n.(string)+"/"+thing[0]+".png"
	//fmt.Println(qc)
	//ctx.String(200,"ok")
	err:=qrcode.WriteFile(url,qrcode.Medium,256,qc)
	if foundation.Wrong(err) {
		foundation.SendMess(ctx,005,"不知道")
		return
	}
	ctx.File(qc)
}