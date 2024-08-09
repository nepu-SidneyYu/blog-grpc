package mysql

import (
	"context"
	"fmt"
	"sync"

	"github.com/nepu-SidneyYu/blog-grpc/internal/config"
	"github.com/nepu-SidneyYu/blog-grpc/internal/logs"
	"github.com/nepu-SidneyYu/blog-grpc/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	_db     *gorm.DB
	_dbOnce sync.Once
)

func Init() {
	_dbOnce.Do(func() {
		var err error
		// 获取配置文件中的MySql配置
		conf := config.GetConfig().MySql

		// 拼接DSN（Data Source Name）
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database, conf.Charset)
		// 使用DSN连接数据库
		_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		// 如果连接失败，则输出错误信息并退出程序
		if err != nil {
			logs.Fatal(context.Background(), "connect mysql failed", zap.String("error", err.Error()))
		}
		// 自动迁移数据库表结构
		err = _db.AutoMigrate(
			&model.UserAuth{},
			&model.Session{},
			&model.Chat{},
		)
		// 如果迁移失败，则输出错误信息并退出程序
		if err != nil {
			logs.Fatal(context.Background(), "migrate mysql failed", zap.String("error", err.Error()))
		}
	})
}
