package server

import (
	proto "github.com/nepu-SidneyYu/blog-grpc/proto/blogFront"
	"google.golang.org/grpc"
)

type Server struct {
}

func (*Server) Servvice() {
	server := grpc.NewServer()
	proto.RegisterCourseServer(server, nil)
}
