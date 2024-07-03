package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/nepu-SidneyYu/blog-grpc/internal/business"
	"github.com/nepu-SidneyYu/blog-grpc/internal/interceptor"
	"github.com/nepu-SidneyYu/blog-grpc/internal/logs"
	v1 "github.com/nepu-SidneyYu/blog-grpc/proto/blog"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (*Server) Service() {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.TraceServerInterceptor()),
	)
	v1.RegisterUserServer(server, &business.UserManager{})
	logs.Info(context.Background(), "服务注册成功")
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", 8080))
	if err != nil {
		logs.Fatal(context.Background(), "failed to listen:")
	}
	go func() {
		err := server.Serve(listener)
		if err != nil {
			logs.Fatal(context.Background(), "failed to start grpc serve:")
		}
	}()
	logs.Info(context.Background(), "grpc server start success")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	select {
	case sig := <-quit:
		logs.Info(context.Background(), "收到关闭信号", zap.String("sig", sig.String()))
	}

	logs.Info(context.Background(), "服务关闭成功...")
	server.GracefulStop()

	logs.Info(context.Background(), "服务停止成功...")
}
