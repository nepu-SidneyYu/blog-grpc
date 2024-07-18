package consts

import (
	"errors"
)

type userErr error

var (
	UserLoginErr                   userErr = errors.New("登陆密码错误")
	UserNotFoundErr                userErr = errors.New("用户不存在")
	UserNameOrPasswordIsNULL       userErr = errors.New("用户名或密码为空")
	UserNameIsNull                 userErr = errors.New("用户名为空")
	UserNameIsExist                userErr = errors.New("用户名已存在")
	GetUserNameFailed              userErr = errors.New("获取用户名失败")
	UserRegisterErr                userErr = errors.New("注册失败")
	UserRegisterPasswordIsNULL     userErr = errors.New("注册密码为空")
	UserRegisterPasswordEncryptErr userErr = errors.New("密码加密失败")
	SendEmailCodeErr               userErr = errors.New("发送邮件验证码失败")
	EmailIsNULL                    userErr = errors.New("邮箱为空")
	EmailCodeErr                   userErr = errors.New("邮箱验证码错误")
	BindEmailErr                   userErr = errors.New("绑定邮箱失败")
	SetCodeErr                     userErr = errors.New("设置验证码失败")
	SendPhoneCodeErr               userErr = errors.New("发送手机验证码失败")
	PhoneIsNULL                    userErr = errors.New("手机号为空")
	//StorgegePhoneCodeErr           userErr = errors.New("存储手机验证码失败")
)

type userErrCode int32

const (
	UserLoginErrCode userErrCode = 1000 + iota // 用户名或密码错误
	UserNameExistErrCode
	UserRegisterErrCode
	SendEmailCodeErrCode
	BindEmailErrCode
	SetCodeErrCode
	SendPhoneCodeErrCode
)

const (
	EmailCodeFeild = "email_code"
	PhoneCodeFeild = "phone_code"
)
