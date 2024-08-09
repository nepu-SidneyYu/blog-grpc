package model

import "gorm.io/gorm"

type Request struct {
	Model       string     `json:"model"`
	Messages    []*Message `json:"messages"`
	ServiceCode string     `json:"serviceCode"`
	MaxTokens   int        `json:"maxTokens"`
	KeyType     string     `json:"keyType"`
	Tag         string     `json:"tag"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	Content string `json:"content"`
	Stop    bool   `json:"stop"`
}

type Session struct {
	// SessionID is the unique identifier for the session
	SessionID string `json:"sessionId" gorm:"column:session_id;primary_key;type:varchar(255);not null"`
	// UserID is the user who initiated the session
	UserID int `json:"userId" gorm:"column:user_id;type:int;not null;index:idx_user_id"`
	// Messages is a slice of messages exchanged during the session
	SessionName string `json:"sessionName" gorm:"column:session_name;type:varchar(255)"`
	// CreatedAt is the timestamp when the session was created
	CreatedAt int64          `gorm:"column:created_at;comment:创建时间;type:bigint(20)" json:"createdAt"`
	UpdatedAt *int64         `gorm:"column:updated_at;comment:更新时间;type:bigint(20)" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间;type:bigint(20)" json:"deletedAt"`
}

type Chat struct {
	ChatId    string         `json:"chatId" gorm:"column:chat_id;primary_key;type:varchar(255);not null"`
	SessionId string         `json:"sessionId" gorm:"column:session_id;type:varchar(255);not null;index:idx_session_id"`
	History   string         `json:"history" gorm:"column:history;type:text"`
	CreatedAt int64          `gorm:"column:created_at;comment:创建时间;type:bigint(20)" json:"createdAt"`
	UpdatedAt *int64         `gorm:"column:updated_at;comment:更新时间;type:bigint(20)" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间;type:bigint(20)" json:"deletedAt"`
}
