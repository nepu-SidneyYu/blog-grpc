package consts

import (
	"errors"
)

type UserErr error

var (
	UserLoginErr             UserErr = errors.New("登陆密码错误")
	UserNotFoundErr          UserErr = errors.New("用户不存在")
	UserNameOrPasswordIsNULL UserErr = errors.New("用户名或密码为空")
)

type UserErrCode int32

const (
	UserLoginErrCode UserErrCode = 1000 + iota // 用户名或密码错误
)
