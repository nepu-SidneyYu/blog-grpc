package repository

import (
	"sync"

	"github.com/nepu-SidneyYu/blog-grpc/internal/repository/mysql"
	"github.com/nepu-SidneyYu/blog-grpc/internal/repository/redis"
)

var (
	_blogUserRepository     User
	_blogUserRepositoryOnce sync.Once // 保证单例
)

func Init() {
	mysql.Init()
	redis.Init()
}

func GetBlogUserRepository() User {
	_blogUserRepositoryOnce.Do(func() {
		_blogUserRepository = mysql.NewUser()
	})
	return _blogUserRepository
}
