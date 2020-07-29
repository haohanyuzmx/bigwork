# Web 后端 暑期考核 网盘

##### 使用介绍

###### POST  /register：注册，（username唯一）

输入

```json
{
    “username”:"xxx",
    "password":"xxx"
}
```

返回

```json
{
	"code":200,
	"mess":"成功"
}
```

###### POST  /login：登录 

输入

```json
{
    “username”:"xxx",
    "password":"xxx"
}
```

返回 mess为token

```json
{
    "code": 200,
    "mess": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJyZWRyb2NrIiwiZXhwIjoiMTU5NjAyMDMxMCIsImlhdCI6IjE1OTYwMDk1MTAiLCJ1c2VybmFtZSI6InRlc3QifQ==.UNO/QmDYTuI2bZteKjs28rVWpgHsuUVWznv7ihYxH4c="
}
```

###### POST /user/up：文件上传

输入 form-data：key为file ，value为文件

返回

```json
{
	"code":200,
	"mess":"成功"
}
```

###### GET  /down/*file：文件下载

###### GET  /user/show：展示你的文件的下载地址

###### POST /user/share：二维码分享

输入 key：url，value：xxx

返回二维码





### 其他

断点续传是靠记录上传的量到redis，第二次上传的时候用文件偏移量继续写入文件

加密的只是url且用的是base64

二维码分享其实只用了一个库的命令

下载限速的话就不能简单使用static而要自己控制写入网页速度就没做了

最后半天还是没做到docker