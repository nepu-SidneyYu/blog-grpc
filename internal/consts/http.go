package consts

type StatusCode = int32

type StatusMessage = string

const (
	StatusOK StatusCode = 0 // 访问成功状态码
)

const (
	StatusSuccess StatusMessage = "success" // 访问成功状态信息
)
