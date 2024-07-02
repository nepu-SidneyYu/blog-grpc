package business

import (
	"github.com/nepu-SidneyYu/blog-grpc/internal/repository"
	"github.com/nepu-SidneyYu/blog-grpc/proto/blog"
)

type UserManager struct {
	blog.UnimplementedUserServer
	userRepository repository.User
}
