package mysql

import "github.com/nepu-SidneyYu/blog-grpc/internal/model"

type Chat struct {
}

func NewChat() *Chat {
	return &Chat{}
}

func SetSession(session *model.Session) error {
	return _db.Model(&model.Session{}).Create(&session).Error
}
func SetChat(chat *model.Chat) error {
	return _db.Model(&model.Chat{}).Create(&chat).Error
}
