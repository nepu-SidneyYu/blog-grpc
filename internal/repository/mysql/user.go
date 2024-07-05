package mysql

import (
	"time"

	"github.com/nepu-SidneyYu/blog-grpc/internal/model"
)

type User struct {
	// 定义User结构体
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetUserByName(name string) (*model.UserAuth, error) {
	var userauth model.UserAuth
	tx := _db.Model(&model.UserAuth{}).Where("username = ?", name).First(&userauth)
	// 根据id查询用户信息
	// 返回一个User结构体和一个错误信息
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &userauth, nil
}
func (u *User) SetUser(username, password string) error {
	tx := _db.Model(&model.UserAuth{}).Create(&model.UserAuth{Username: username, Password: password, CreatedAt: time.Now().UnixMilli()})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
