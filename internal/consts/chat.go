package consts

import "errors"

const (
	ModelName   = "gpt-4"
	ServiceCode = "SC_MEDIA_CLOUD"
	Tag         = "panorama"
	KeyType     = "Z120"
	MaxTokens   = 2048
)

const (
	ChatHistoryFeild = "chat_history"
	ChatMessageFeild = "chat_message"
)

type ChatErr error

var (
	CreateSessionErr ChatErr = errors.New("create session err")
	SessionIdIsNULL  ChatErr = errors.New("session id is null")
)

type ChatErrCode int32

const (
	CreateSessionErrCode ChatErrCode = 1100 + iota
)
