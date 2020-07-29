package module

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	var u User
	err:=DB.Where(&User{
		Username: "test",
		Password: "23",
	}).First(&u).Error
	fmt.Println(u,err)
}
func TestShowMyFile(t *testing.T) {
	//conn:=Pool.Get()
	//conn.Do("set","test/1.png",198)
	//x,err:=redis.Int64(conn.Do("get","tet/1.png"))
	//fmt.Println(x,err)
	err:=DB.Create(&MyFile{
		Username: "test",
		Uri: "localhost:8080/down/test/1.txt",
	}).Error
	fmt.Println(err)
}