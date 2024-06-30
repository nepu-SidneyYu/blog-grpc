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
	// 定义一个全局变量_cacheOnce，用于控制redis连接的初始化
	_cacheOnce.Do(func() {
		// 从配置文件中获取Redis配置
		conf := config.GetConfig().Redis
		// 创建一个redis客户端
		_cache = redis.NewClient(&redis.Options{
			// 设置redis地址
			Addr: fmt.Sprintf("%v:%v", conf.Host, conf.Port),
			// 设置redis密码
			Password: conf.Password,
			// 设置redis数据库
			DB: conf.DB,
		})
		// 尝试与redis服务器建立连接
		err := _cache.Ping(context.Background()).Err()
		// 如果连接失败，则输出错误信息并退出程序
		if err != nil {
			logs.Fatal(context.Background(), "connect redis failed", zap.String("error", err.Error()))
		}
		logs.Info(context.Background(), "connect redis success", zap.String("Info", "connect redis success"))
	})
}
