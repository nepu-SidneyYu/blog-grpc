package mysql

import "github.com/nepu-SidneyYu/blog-grpc/internal/model"

type User struct {
	// 定义User结构体
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetUserByName(name string) (*model.UserAuth, error) {
	_db.Model(&model.UserAuth{})
	// 根据id查询用户信息
	// 返回一个User结构体和一个错误信息
	return nil, nil
}
