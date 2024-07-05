package business

import (
	"context"
	"errors"
	"fmt"

	"github.com/nepu-SidneyYu/blog-grpc/internal/config"
	"github.com/nepu-SidneyYu/blog-grpc/internal/consts"
	"github.com/nepu-SidneyYu/blog-grpc/internal/logs"
	"github.com/nepu-SidneyYu/blog-grpc/internal/repository"
	"github.com/nepu-SidneyYu/blog-grpc/internal/utils"
	"github.com/nepu-SidneyYu/blog-grpc/internal/utils/jwt"
	blog "github.com/nepu-SidneyYu/blog-grpc/proto/blog"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserManager struct {
	blog.UnimplementedUserServer
	userRepository repository.User
}

func NewUserManager() *UserManager {
	return &UserManager{
		userRepository: repository.GetBlogUserRepository(),
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
	ipSourse := utils.GetIpSource(ip)
	fmt.Println("ipSourse:", ipSourse)
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
	fmt.Println(token)
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

func (u *UserManager) UserRegister(ctx context.Context, req *blog.UserRegisterRequest) (*blog.EmptyResponse, error) {
	if req.Password == "" {
		logs.Error(ctx, "密码为空", zap.String("Error", consts.UserNameOrPasswordIsNULL.Error()))
		return newEmptyResponse(withEmptyResponse(consts.StatusOK, consts.UserRegisterPasswordIsNULL.Error())), nil
	}
	hashpassword, err := utils.BcryptHash(req.Password)
	if err != nil {
		logs.Error(ctx, "密码加密失败", zap.String("Error", err.Error()))
		return newEmptyResponse(withEmptyResponse(int32(consts.UserRegisterErrCode), consts.UserRegisterPasswordEncryptErr.Error())), nil
	}
	err = u.userRepository.SetUser(req.Username, hashpassword)
	if err != nil {
		logs.Error(ctx, "注册用户失败", zap.String("Error", err.Error()))
		return newEmptyResponse(withEmptyResponse(consts.StatusOK, consts.UserRegisterErr.Error())), nil
	}
	return newEmptyResponse(), nil
}
