package business

import (
	"context"
	"errors"

	"github.com/nepu-SidneyYu/blog-grpc/internal/config"
	"github.com/nepu-SidneyYu/blog-grpc/internal/consts"
	"github.com/nepu-SidneyYu/blog-grpc/internal/logs"
	"github.com/nepu-SidneyYu/blog-grpc/internal/repository"
	"github.com/nepu-SidneyYu/blog-grpc/internal/utils"
	"github.com/nepu-SidneyYu/blog-grpc/internal/utils/jwt"
	"github.com/nepu-SidneyYu/blog-grpc/internal/utils/sendcode"
	blog "github.com/nepu-SidneyYu/blog-grpc/proto/blog"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserManager struct {
	blog.UnimplementedUserServer
	userRepository      repository.User
	codeCacheRepository repository.CodeCache
	conf                config.SendEmailCodeConfig
}

func NewUserManager() *UserManager {
	return &UserManager{
		userRepository:      repository.GetBlogUserRepository(),
		codeCacheRepository: repository.GetBlogCodeCacheRepository(),
		conf:                config.GetConfig().Email,
	}
}

type (
	UserLoginResponseFun     func(c *blog.UserLoginResponse)
	UserNameExistResponseFun func(c *blog.UserNameExistResponse)
	EmptyResponseFun         func(c *blog.EmptyResponse)
)

func newUserLoginResponse(s ...UserLoginResponseFun) *blog.UserLoginResponse {
	cp := &blog.UserLoginResponse{
		Code: int32(consts.UserLoginErrCode),
		Msg:  consts.UserLoginErr.Error(),
		Data: nil,
	}
	for _, o := range s {
		o(cp)
	}
	return cp
}
func withUserLoginResponse(code int32, msg string, data *blog.UserInfo) UserLoginResponseFun {
	return func(b *blog.UserLoginResponse) {
		b.Code = code
		b.Msg = msg
		b.Data = data
	}
}
func newUserNameExistResponse(s ...UserNameExistResponseFun) *blog.UserNameExistResponse {
	cp := &blog.UserNameExistResponse{
		Code: int32(consts.UserNameExistErrCode),
		Msg:  consts.UserNameIsExist.Error(),
		Data: nil,
	}
	for _, o := range s {
		o(cp)
	}
	return cp
}
func withUserNameExistResponse(code int32, msg string, data *blog.UserNameExistInfo) UserNameExistResponseFun {
	return func(b *blog.UserNameExistResponse) {
		b.Code = code
		b.Msg = msg
		b.Data = data
	}
}
func newEmptyResponse(s ...EmptyResponseFun) *blog.EmptyResponse {
	cp := &blog.EmptyResponse{
		Code: consts.StatusOK,
		Msg:  consts.StatusSuccess,
		Data: nil,
	}
	for _, o := range s {
		o(cp)
	}
	return cp
}
func withEmptyResponse(code int32, msg string) EmptyResponseFun {
	return func(b *blog.EmptyResponse) {
		b.Code = code
		b.Msg = msg
	}
}

func (u *UserManager) UserLogin(ctx context.Context, req *blog.UserLoginRequest) (*blog.UserLoginResponse, error) {
	if req.Username == "" || req.Password == "" {
		logs.Error(ctx, "用户名或密码为空", zap.String("Error", consts.UserNameOrPasswordIsNULL.Error()))
		return newUserLoginResponse(withUserLoginResponse(int32(consts.UserLoginErrCode), consts.UserNameOrPasswordIsNULL.Error(), nil)), nil
	}
	ip, ok := ctx.Value("ip").(string)
	if !ok || ip == "" {
		logs.Error(ctx, "获取ip失败", zap.String("Error", "从ctx中获取ip失败"))
		ip = "127.0.0.1"
	}
	//ipSourse := utils.GetIpSource(ip)

	// 验证用户名和密码
	userInfo, err := u.userRepository.GetUserByName(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error(ctx, "用户不存在", zap.String("Error", err.Error()))
			return newUserLoginResponse(withUserLoginResponse(int32(consts.UserLoginErrCode), consts.UserNotFoundErr.Error(), nil)), nil
		}
		logs.Error(ctx, "查询用户失败", zap.String("Error", err.Error()))
		return newUserLoginResponse(withUserLoginResponse(int32(consts.UserLoginErrCode), consts.UserLoginErr.Error(), nil)), nil
	}
	if !utils.BcryptCheck(req.Password, userInfo.Password) {
		logs.Error(ctx, "密码错误", zap.String("Error", consts.UserLoginErr.Error()))
		return newUserLoginResponse(withUserLoginResponse(int32(consts.UserLoginErrCode), consts.UserLoginErr.Error(), nil)), nil
	}
	// 生成token
	conf := config.GetConfig().JWt
	token, err := jwt.CreateToken(conf.Secret, conf.Issuer, int(conf.Expire), userInfo.ID)
	if err != nil {
		logs.Error(ctx, "生成token失败", zap.String("Error", err.Error()))
		return newUserLoginResponse(withUserLoginResponse(int32(consts.UserLoginErrCode), consts.UserLoginErr.Error(), nil)), nil
	}
	data := &blog.UserInfo{
		Id:              int32(userInfo.ID),
		Nickname:        "测试用户",
		Token:           token,
		ArticlesLikeSet: []string{},
		CommentsLikeSet: []string{},
	}
	logs.Info(ctx, "登录成功", zap.String("UserName", req.Username))
	return newUserLoginResponse(withUserLoginResponse(consts.StatusOK, consts.StatusSuccess, data)), nil
}

func (u *UserManager) UserNameExist(ctx context.Context, req *blog.UserNameExistRequest) (*blog.UserNameExistResponse, error) {
	if req.Username == "" {
		logs.Error(ctx, "用户名为空", zap.String("Error", consts.UserNameOrPasswordIsNULL.Error()))
		return newUserNameExistResponse(withUserNameExistResponse(int32(consts.UserNameExistErrCode), consts.UserNameIsNull.Error(), nil)), nil
	}
	_, err := u.userRepository.GetUserByName(req.Username)
	if err == nil {
		logs.Error(ctx, "用户名已存在", zap.String("Error", consts.UserNameIsExist.Error()))
		return newUserNameExistResponse(withUserNameExistResponse(int32(consts.UserNameExistErrCode), consts.UserNameIsExist.Error(), nil)), nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		logs.Error(ctx, "查询用户失败", zap.String("Error", err.Error()))
		return newUserNameExistResponse(withUserNameExistResponse(int32(consts.UserNameExistErrCode), consts.GetUserNameFailed.Error(), nil)), nil
	}
	return newUserNameExistResponse(withUserNameExistResponse(consts.StatusOK, consts.StatusSuccess, &blog.UserNameExistInfo{Exist: false})), nil
}

func (u *UserManager) SendPhoneCode(ctx context.Context, req *blog.SendPhoneCodeRequest) (*blog.EmptyResponse, error) {
	if req.Phone == "" {
		logs.Error(ctx, "手机号为空", zap.String("Error", consts.PhoneIsNULL.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.SendPhoneCodeErrCode), consts.PhoneIsNULL.Error())), nil
	}
	logs.Info(ctx, "发送短信验证码", zap.String("Phone", req.Phone))
	//发送短信验证码
	code, err := sendcode.SendPhoneCode(req.Phone)
	logs.Info(ctx, "发送短信验证码", zap.String("Code", code))
	if err != nil {
		logs.Error(ctx, "发送短信验证码失败", zap.String("Error", err.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.SendPhoneCodeErrCode), consts.SendPhoneCodeErr.Error())), nil
	}
	err = u.codeCacheRepository.SetPhoneCode(consts.PhoneCodeFeild, req.Phone, code, int64(config.GetConfig().Phone.Expire))
	if err != nil {
		logs.Error(ctx, "保存验证码失败", zap.String("Error", err.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.SendPhoneCodeErrCode), consts.SetCodeErr.Error())), nil
	}
	return newEmptyResponse(), nil
}

func (u *UserManager) UserRegister(ctx context.Context, req *blog.UserRegisterRequest) (*blog.EmptyResponse, error) {
	if req.Phone == "" {
		logs.Error(ctx, "手机号为空", zap.String("Error", consts.PhoneIsNULL.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.UserRegisterErrCode), consts.PhoneIsNULL.Error())), nil
	}
	if req.Password == "" {
		logs.Error(ctx, "密码为空", zap.String("Error", consts.UserNameOrPasswordIsNULL.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.UserRegisterErrCode), consts.UserRegisterPasswordIsNULL.Error())), nil
	}
	if req.Code == "" {
		logs.Error(ctx, "验证码为空", zap.String("Error", consts.CodeIsNULL.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.UserRegisterErrCode), consts.CodeIsNULL.Error())), nil
	}
	code, err := u.codeCacheRepository.GetPhoneCode(consts.PhoneCodeFeild, req.Phone)
	if err != nil {
		logs.Error(ctx, "获取验证码失败", zap.String("Error", err.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.UserRegisterErrCode), consts.GetPhoneCodeErr.Error())), nil
	}
	if code != req.Code {
		logs.Error(ctx, "验证码错误", zap.String("Error", "验证码错误"))
		return newEmptyResponse(withEmptyResponse(int32(consts.UserRegisterErrCode), consts.CodeIsErr.Error())), nil
	}
	hashpassword, err := utils.BcryptHash(req.Password)
	if err != nil {
		logs.Error(ctx, "密码加密失败", zap.String("Error", err.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.UserRegisterErrCode), consts.UserRegisterPasswordEncryptErr.Error())), nil
	}
	err = u.userRepository.SetUser(req.Phone, hashpassword)
	if err != nil {
		logs.Error(ctx, "注册用户失败", zap.String("Error", err.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.UserRegisterErrCode), consts.UserRegisterErr.Error())), nil
	}
	return newEmptyResponse(), nil
}

// 绑定邮箱
func (u *UserManager) BindEmail(ctx context.Context, req *blog.BindEmailRequest) (*blog.EmptyResponse, error) {
	if req.Email == "" {
		logs.Error(ctx, "邮箱为空", zap.String("Error", consts.EmailIsNULL.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.BindEmailErrCode), consts.EmailIsNULL.Error())), nil
	}
	//TODO：验证验证码是否正确
	code, err := u.codeCacheRepository.GetEmailCode(consts.EmailCodeFeild, req.Email)
	if err != nil {
		logs.Error(ctx, "获取验证码失败", zap.String("Error", err.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.BindEmailErrCode), "获取验证码失败")), nil
	}
	if code != req.Code {
		logs.Error(ctx, "验证码错误", zap.String("Error", "验证码错误"))
		return newEmptyResponse(withEmptyResponse(int32(consts.BindEmailErrCode), "验证码错误")), nil
	}
	//TODO：绑定邮箱
	err = u.userRepository.BindEmail(req.Phone, req.Email)
	if err != nil {
		logs.Error(ctx, "绑定邮箱失败", zap.String("Error", err.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.BindEmailErrCode), consts.BindEmailErr.Error())), nil
	}
	return newEmptyResponse(), nil
}

// 发送邮箱验证码
func (u *UserManager) SendEmailCode(ctx context.Context, req *blog.SendCodeRequest) (*blog.EmptyResponse, error) {
	if req.Email == "" {
		logs.Error(ctx, "邮箱为空", zap.String("Error", "邮箱为空"))
		return newEmptyResponse(withEmptyResponse(int32(consts.SendEmailCodeErrCode), consts.EmailIsNULL.Error())), nil
	}
	code, err := sendcode.SendVerificationCode(ctx, req.Email)
	if err != nil {
		logs.Error(ctx, "发送验证码失败", zap.String("Error", err.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.SendEmailCodeErrCode), consts.SendEmailCodeErr.Error())), nil
	}
	//TODO:设置验证码的过期时间
	err = u.codeCacheRepository.SetEmailCode(consts.EmailCodeFeild, req.Email, code, int64(u.conf.Expire))
	if err != nil {
		logs.Error(ctx, "存储验证码失败", zap.String("Error", err.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.SetCodeErrCode), consts.SetCodeErr.Error())), nil
	}
	return newEmptyResponse(), nil
}

func (u *UserManager) SetUserName(ctx context.Context, req *blog.SetUserNameRequest) (*blog.EmptyResponse, error) {
	if req.Username == "" {
		logs.Error(ctx, "用户名为空", zap.String("Error", consts.UserNameIsNull.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.SetUserNameErrCode), consts.UserNameIsNull.Error())), nil
	}
	err := u.userRepository.SetUserName(req.Phone, req.Username)
	if err != nil {
		logs.Error(ctx, "设置用户名失败", zap.String("Error", err.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.SetUserNameErrCode), consts.SerUserNameErr.Error())), nil
	}
	return newEmptyResponse(), nil
}

func (u *UserManager) UserUsePhoneLogin(ctx context.Context, req *blog.UserUsePhoneLoginRequest) (*blog.UserLoginResponse, error) {
	if req.Phone == "" {
		logs.Error(ctx, "手机号为空", zap.String("Error", consts.PhoneIsNULL.Error()))
		return newUserLoginResponse(withUserLoginResponse(int32(consts.UserUsePhoneLoginErrCode), consts.PhoneIsNULL.Error(), nil)), nil
	}
	user, err := u.userRepository.GetUserByPhone(req.Phone)
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			logs.Error(ctx, "用户不存在", zap.String("Error", err.Error()))
			return newUserLoginResponse(withUserLoginResponse(int32(consts.UserUsePhoneLoginErrCode), consts.PhoneNotRegister.Error(), nil)), nil
		}
		logs.Error(ctx, "获取用户失败", zap.String("Error", err.Error()))
		return newUserLoginResponse(withUserLoginResponse(int32(consts.UserUsePhoneLoginErrCode), "获取用户信息失败", nil)), nil
	}
	if !utils.BcryptCheck(req.Password, user.Password) {
		logs.Error(ctx, "密码错误", zap.String("Error", consts.UserLoginErr.Error()))
		return newUserLoginResponse(withUserLoginResponse(int32(consts.UserUsePhoneLoginErrCode), consts.UserLoginErr.Error(), nil)), nil
	}
	// 生成token
	conf := config.GetConfig().JWt
	token, err := jwt.CreateToken(conf.Secret, conf.Issuer, int(conf.Expire), user.ID)
	if err != nil {
		logs.Error(ctx, "生成token失败", zap.String("Error", err.Error()))
		return newUserLoginResponse(withUserLoginResponse(int32(consts.UserUsePhoneLoginErrCode), consts.UserLoginErr.Error(), nil)), nil
	}
	data := &blog.UserInfo{
		Id:              int32(user.ID),
		Nickname:        "测试用户",
		Token:           token,
		ArticlesLikeSet: []string{},
		CommentsLikeSet: []string{},
	}
	logs.Info(ctx, "登录成功", zap.String("UserName", req.Phone))
	return newUserLoginResponse(withUserLoginResponse(consts.StatusOK, consts.StatusSuccess, data)), nil
}
func (u *UserManager) UserUseEmailLogin(ctx context.Context, req *blog.UserUseEmailLoginRequest) (*blog.UserLoginResponse, error) {
	if req.Email == "" {
		logs.Error(ctx, "邮箱为空", zap.String("Error", consts.EmailIsNULL.Error()))
		return newUserLoginResponse(withUserLoginResponse(int32(consts.UserUseEmailLoginErrCode), consts.EmailIsNULL.Error(), nil)), nil
	}
	user, err := u.userRepository.GetUserByEmail(req.Email)
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			logs.Error(ctx, "用户不存在", zap.String("Error", err.Error()))
			return newUserLoginResponse(withUserLoginResponse(int32(consts.UserUseEmailLoginErrCode), consts.EmailNotBind.Error(), nil)), nil
		}
		logs.Error(ctx, "获取用户失败", zap.String("Error", err.Error()))
		return newUserLoginResponse(withUserLoginResponse(int32(consts.UserUseEmailLoginErrCode), "获取用户信息失败", nil)), nil
	}
	if !utils.BcryptCheck(req.Password, user.Password) {
		logs.Error(ctx, "密码错误", zap.String("Error", consts.UserLoginErr.Error()))
		return newUserLoginResponse(withUserLoginResponse(int32(consts.UserUseEmailLoginErrCode), consts.UserLoginErr.Error(), nil)), nil
	}
	// 生成token
	conf := config.GetConfig().JWt
	token, err := jwt.CreateToken(conf.Secret, conf.Issuer, int(conf.Expire), user.ID)
	if err != nil {
		logs.Error(ctx, "生成token失败", zap.String("Error", err.Error()))
		return newUserLoginResponse(withUserLoginResponse(int32(consts.UserUseEmailLoginErrCode), consts.UserLoginErr.Error(), nil)), nil
	}
	data := &blog.UserInfo{
		Id:              int32(user.ID),
		Nickname:        "测试用户",
		Token:           token,
		ArticlesLikeSet: []string{},
		CommentsLikeSet: []string{},
	}
	logs.Info(ctx, "登录成功", zap.String("UserEmail", req.Email))
	return newUserLoginResponse(withUserLoginResponse(consts.StatusOK, consts.StatusSuccess, data)), nil
}
