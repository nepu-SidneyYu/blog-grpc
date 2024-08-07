package business

import (
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
	// for i := 0; i < 500; i++ {
	// 	time.Sleep(500 * time.Millisecond)
	// 	if err := stream.Send(&blog.ChatResponse{Code: consts.StatusOK, Msg: consts.StatusSuccess, Content: strconv.Itoa(i + 1)}); err != nil {
	// 		logs.Error(stream.Context(), "stream send error", zap.String("error", err.Error()))
	// 	}
	// }
	//
	done := make(chan bool)
	go utils.Chat(request, func(r *model.Response) {
		if err := stream.Send(&blog.ChatResponse{Code: consts.StatusOK, Msg: consts.StatusSuccess, Content: r.Content}); err != nil {
			logs.Error(stream.Context(), "stream send error", zap.String("error", err.Error()))
		}
		if r.Stop {
			done <- true
		}
	})
	<-done

	// go utils.Chat(request, func(r *model.Response) {
	// 	if err := stream.Send(&blog.ChatResponse{Code: consts.StatusOK, Msg: consts.StatusSuccess, Content: r.Content}); err != nil {
	// 		logs.Error(stream.Context(), "stream send error", zap.String("error", err.Error()))
	// 	}
	// })

	// utils.Chat(request, func(r *model.Response) {
	// 	if err := stream.Send(&blog.ChatResponse{Code: consts.StatusOK, Msg: consts.StatusSuccess, Content: r.Content}); err != nil {
	// 		logs.Error(context.Background(), "stream send error", zap.String("error", err.Error()))
	// 	}
	// 	//time.Sleep(1 * time.Second)
	// })
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
