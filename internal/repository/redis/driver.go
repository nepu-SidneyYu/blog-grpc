package redis

import (
	"context"
	"fmt"
	"sync"

	"github.com/nepu-SidneyYu/blog-grpc/internal/config"
	"github.com/nepu-SidneyYu/blog-grpc/internal/logs"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var (
	_cache     *redis.Client
	_cacheOnce sync.Once
)

func Init() {
	_cacheOnce.Do(func() {
		conf := config.GetConfig().Redis
		_cache = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%v:%v", conf.Host, conf.Port),
			Password: conf.Password,
			DB:       conf.DB,
		})
		err := _cache.Ping(context.Background()).Err()
		if err != nil {
			logs.Fatal(context.Background(), "connect redis failed", zap.String("error", err.Error()))
		}
	})
}
