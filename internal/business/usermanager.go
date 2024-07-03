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
	UserLoginResponseFun func(c *blog.UserLoginResponse)
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

func (u *UserManager) UserLogin(ctx context.Context, req *blog.UserLoginRequest) (*blog.UserLoginResponse, error) {
	if req.Username == "" || req.Password == "" {
		logs.Error(ctx, "用户名或密码为空", zap.String("Error", consts.UserNameOrPasswordIsNULL.Error()))
		return newUserLoginResponse(withUserLoginResponse(int32(consts.UserLoginErrCode), consts.UserNameOrPasswordIsNULL.Error(), nil)), nil
	}
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
