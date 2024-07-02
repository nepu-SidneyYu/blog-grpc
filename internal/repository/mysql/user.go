package mysql

import "github.com/nepu-SidneyYu/blog-grpc/internal/model"

type User struct {
	// 定义User结构体
}

func (u *User) NewUser() *User {
	return &User{}
}

func (u *User) GetUserById(id int) (*model.UserAuth, error) {
	// 根据id查询用户信息
	// 返回一个User结构体和一个错误信息
	return nil, nil
}
