package module
import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type User struct {
	gorm.Model
	Username string `json:"username"gorm:"unique_index"`
	Password string `json:"password"`
}
type MyFile struct {
	gorm.Model
	Username string
	Uri string
}
var Pool redis.Pool
var DB *gorm.DB
func init()  {
	var err error
	db, err := gorm.Open("mysql","mysql","root:@tcp(127.0.0.1:3306)/summer_work?parseTime=true&charset=utf8&loc=Local")
	if err != nil {
		log.Panicln(err)
	}
	DB = db
	Pool = redis.Pool{
		MaxIdle:     16,
		MaxActive:   32,
		IdleTimeout: 120,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	if Pool.Get()==nil {
		log.Println("redis连接失败")
		return
	}
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
	if !db.HasTable(&MyFile{}) {
		db.CreateTable(&MyFile{})
	}
}
