package business

import (
	"context"
	"strconv"
	"time"

	"github.com/nepu-SidneyYu/blog-grpc/internal/consts"
	blog "github.com/nepu-SidneyYu/blog-grpc/proto/blog"
)

type ChatManager struct {
	blog.UnimplementedChatServer
}

func NewChatManager() *ChatManager { return &ChatManager{} }

func (c *ChatManager) Chat(ctx context.Context, req *blog.ChatRequest, stream blog.Chat_ChatServer) (*blog.ChatResponse, error) {

	for i := 0; i < 5; i++ { // 假设我们发送5个部分响应
		time.Sleep(1 * time.Second) // 模拟处理时间
		if err := stream.Send(&blog.ChatResponse{Code: consts.StatusOK, Msg: consts.StatusSuccess, Content: strconv.Itoa(i + 1)}); err != nil {
			return nil, err
		}
	}
	return nil, nil
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
