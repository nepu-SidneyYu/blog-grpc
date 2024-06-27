package repository

import (
	"github.com/nepu-SidneyYu/blog-grpc/internal/repository/mysql"
	"github.com/nepu-SidneyYu/blog-grpc/internal/repository/redis"
)

var ()

func Init() {
	mysql.Init()
	redis.Init()
}
