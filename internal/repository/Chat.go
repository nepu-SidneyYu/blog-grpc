package repository

import "github.com/nepu-SidneyYu/blog-grpc/internal/model"

type Chat interface {
	SetSession(session *model.Session) error
	SetChat(chat *model.Chat) error
}
