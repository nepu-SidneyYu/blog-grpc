package repository

import "github.com/nepu-SidneyYu/blog-grpc/internal/model"

type User interface {
	// 定义用户相关的操作
	GetUserById(id int) (*model.UserAuth, error)
}
