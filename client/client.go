package main

import (
	"context"
	"fmt"
	"io"

	//"github.com/nacos-group/nacos-sdk-go/v2/api/grpc"
	"github.com/nepu-SidneyYu/blog-grpc/internal/interceptor"
	"github.com/nepu-SidneyYu/blog-grpc/internal/logs"
	blog "github.com/nepu-SidneyYu/blog-grpc/proto/blog"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := "192.168.184.129:6666"
	//opts:= make([]grpc.DialOption,0)
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptor.TraceClientInterceptor()),
	)
	if err != nil {
		logs.Fatal(context.Background(), "failed to dial server", zap.String("err", err.Error()))
	}
	defer conn.Close()
	clientChat := blog.NewChatClient(conn)
	//userclient := blog.NewUserClient(conn)
	go chatresponse(clientChat, &blog.ChatRequest{Content: "写一篇800字的关于母亲的作文"})
	//go chatresponse(clientChat, &blog.ChatRequest{Content: "蜀道难"}, 2)
	//time.Sleep(time.Second * 5)
	//go userresponse(userclient, &blog.UserRegisterRequest{Phone: "123456789", Password: "123456", Code: "000000"})
	for {

	}
}

func chatresponse(client blog.ChatClient, req *blog.ChatRequest) {
	stream, err := client.Chat(context.Background(), req)
	if err != nil {
		logs.Error(context.Background(), "failed to chat", zap.String("err", err.Error()))
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			logs.Error(context.Background(), "failed to receive message", zap.String("err", err.Error()))
			break
		}
		if err != nil {
			logs.Error(context.Background(), "failed to receive message", zap.String("err", err.Error()))
		}
		logs.Info(context.Background(), "receive message", zap.String("message", resp.Content))
	}
}

func userresponse(client blog.UserClient, req *blog.UserRegisterRequest) {
	resp, err := client.UserRegister(context.Background(), req)
	if err != nil {
		logs.Error(context.Background(), "failed to register", zap.String("err", err.Error()))
	}
	logs.Info(context.Background(), "register success", zap.String("message", fmt.Sprintf("%#v\n", resp)))
}
