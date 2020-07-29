package module

import (
	"errors"
)

func Register(u User) error {
	err:=DB.Table("users").Where("username=?",u.Username).Error
	if err==nil {
		return errors.New("用户存在")
	}
	DB.Create(&u)
	return nil
}

func Login(u User) error {
	err:=DB.Where(&User{
		Username: u.Username,
		Password: u.Password,
	}).First(&u).Error
	return err
}
