package repository

import "github.com/nepu-SidneyYu/blog-grpc/internal/model"

type User interface {
	// 定义用户相关的操作
	GetUserByName(name string) (*model.UserAuth, error)
	SetUser(Phone, password string) error
	SetUserName(phone, name string) error
	BindEmail(phone, email string) error
	GetUserByPhone(phone string) (*model.UserAuth, error)
	GetUserByEmail(email string) (*model.UserAuth, error)
}
