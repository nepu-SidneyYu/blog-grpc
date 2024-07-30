package business

import (
	"context"

	"github.com/nepu-SidneyYu/blog-grpc/internal/consts"
	"github.com/nepu-SidneyYu/blog-grpc/internal/logs"
	"github.com/nepu-SidneyYu/blog-grpc/internal/model"
	"github.com/nepu-SidneyYu/blog-grpc/internal/utils"
	blog "github.com/nepu-SidneyYu/blog-grpc/proto/blog"
	"go.uber.org/zap"
)

type ChatManager struct {
	blog.UnimplementedChatServer
}

func NewChatManager() *ChatManager { return &ChatManager{} }

func (c *ChatManager) Chat(req *blog.ChatRequest, stream blog.Chat_ChatServer) error {
	// ctx := stream.Context()
	// if ctx.Err() != nil {
	// 	logs.Error(ctx, "context error", zap.String("error", ctx.Err().Error()))
	// 	return ctx.Err()
	// }
	request := &model.Request{
		Model:       consts.ModelName,
		MaxTokens:   consts.MaxTokens,
		ServiceCode: consts.ServiceCode,
		KeyType:     consts.KeyType,
		Tag:         consts.Tag,
	}
	request.Messages = make([]*model.Message, 0)
	request.Messages = append(request.Messages, &model.Message{
		Role:    "user",
		Content: req.Content,
	})
	go func() {
		done := make(chan struct{})
		utils.Chat(request, func(r *model.Response) {
			if err := stream.Send(&blog.ChatResponse{Code: consts.StatusOK, Msg: consts.StatusSuccess, Content: r.Content}); err != nil {
				logs.Error(context.Background(), "stream send error", zap.String("error", err.Error()))
			}
			if r.Stop {
				done <- struct{}{}
			}
		})
		<-done
	}()
	return nil
}

// package main

// import (
//     "context"
//     "log"
//     "time"

//     "google.golang.org/grpc"
//     pb "path/to/your/protobuf/package" // 替换为你的 protobuf 生成的 Go 包的路径
// )

// // server 是 gRPC 服务的实现
// type server struct {
//     pb.UnimplementedChatGPTServiceServer
// }

// // AskStream 方法实现了 ChatGPTService 的 AskStream RPC
// func (s *server) AskStream(req *pb.ChatGPTRequest, stream pb.ChatGPTService_AskStreamServer) error {
//     // 模拟 ChatGPT 的流式响应
//     for i := 0; i < 5; i++ { // 假设我们发送5个部分响应
//         time.Sleep(1 * time.Second) // 模拟处理时间
//         if err := stream.Send(&pb.ChatGPTResponse{Message: "Partial response " + string(i+1)}); err != nil {
//             return err
//         }
//     }
//     return nil
// }

// func main() {
//     lis, err := net.Listen("tcp", ":50051")
//     if err != nil {
//         log.Fatalf("failed to listen: %v", err)
//     }
//     s := grpc.NewServer()
//     pb.RegisterChatGPTServiceServer(s, &server{})
//     if err := s.Serve(lis); err != nil {
//         log.Fatalf("failed to serve: %v", err)
//     }
// }
