package main

import (
	"github.com/nepu-SidneyYu/blog-grpc/internal/config"
	"github.com/nepu-SidneyYu/blog-grpc/internal/loger"
	"github.com/nepu-SidneyYu/blog-grpc/internal/logs"
	"github.com/nepu-SidneyYu/blog-grpc/internal/repository"
	"github.com/nepu-SidneyYu/blog-grpc/internal/server"
)

func main() {
	//日志初始化
	loger.Init()
	logs.SetLogerWriter(loger.Writer())

	//读取配置文件
	config.LoadConfig()

	//数据库初始化
	repository.Init()

	//启动服务
	ser := server.NewServer()
	ser.Service()
}
