package foundation

import (
	"bigwork/module"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func Wrong(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}
func SendMess(ctx *gin.Context, state int, mess string) {
	ctx.JSON(200, gin.H{
		"code": state,
		"mess": mess,
	})
}
func Savefile(file *multipart.FileHeader, dst string) error {
	sf, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer sf.Close()
	if err != nil {
		return err
	}
	conn := module.Pool.Get()
	star, _ := redis.Int64(conn.Do("get", dst))
	f, err := file.Open()
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Seek(star, io.SeekStart)
	if err != nil {
		return err
	}
	buf := make([]byte, 8192)
	for {
		n, err := f.Read(buf)
		if err==io.EOF {
			return nil
		}else if err!=nil {
			return err
		}
		_, err = sf.WriteAt(buf, star)
		if err != nil {
			return err
		}
		star+=int64(n)
		conn.Do("set",dst,star)
	}
}
