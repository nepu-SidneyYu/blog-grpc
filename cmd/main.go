package main

import (
	"fmt"

	"github.com/nepu-SidneyYu/blog-grpc/internal/config"
	"github.com/nepu-SidneyYu/blog-grpc/internal/loger"
	"github.com/nepu-SidneyYu/blog-grpc/internal/logs"
)

func main() {
	//日志初始化
	loger.Init()
	logs.SetLogerWriter(loger.Writer())

	//读取配置文件
	config.LoadConfig()
	conf := config.GetConfig()
	fmt.Println(conf)
	//启动服务
	return
}
