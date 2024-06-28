package mysql

import (
	"context"
	"fmt"
	"sync"

	"github.com/nepu-SidneyYu/blog-grpc/internal/config"
	"github.com/nepu-SidneyYu/blog-grpc/internal/logs"
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
		conf := config.GetConfig().MySql

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database, conf.Charset)
		_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			logs.Fatal(context.Background(), "connect mysql failed", zap.String("error", err.Error()))
		}
		err = _db.AutoMigrate(
		// &mysql.User{},
		)
		if err != nil {
			logs.Fatal(context.Background(), "migrate mysql failed", zap.String("error", err.Error()))
		}
	})
}
