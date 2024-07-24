package repository

import (
	"sync"

	"github.com/nepu-SidneyYu/blog-grpc/internal/repository/mysql"
	"github.com/nepu-SidneyYu/blog-grpc/internal/repository/redis"
)

var (
	_blogUserRepository     User
	_blogUserRepositoryOnce sync.Once // 保证单例

	_blogCodeCacheRepository     CodeCache
	_blogCodeCacheRepositoryOnce sync.Once // 保证单例

	_blogUserNameCacheRepository     UserNameCache
	_blogUserNameCacheRepositoryOnce sync.Once // 保证单例
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

func GetBlogCodeCacheRepository() CodeCache {

	_blogCodeCacheRepositoryOnce.Do(func() {
		_blogCodeCacheRepository = redis.NewCodeCache()
	})
	return _blogCodeCacheRepository
}

func GetUserNameCacheRepository() UserNameCache {
	_blogUserNameCacheRepositoryOnce.Do(func() {
		_blogUserNameCacheRepository = redis.NewUserNameCache()
	})
	return _blogUserNameCacheRepository
}
