package control

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"os"
	"testing"
)

func TestUp(t *testing.T) {
	out,err:=os.Create("../file/name.txt")
	if err!=nil {
		fmt.Println(err)
		return
	}
	a,b:=out.Write([]byte("123"))
	fmt.Println(a,b)
}
func TestShare(t *testing.T) {
	err:=qrcode.WriteFile("localhost/down/1.txt",qrcode.Medium,256,"../share/1.png")
	fmt.Println(err)
}