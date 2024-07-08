package sendcode

import (
	"context"
	"fmt"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
	"github.com/nepu-SidneyYu/blog-grpc/internal/config"
)

// SendVerificationCode sends a verification code to the user's email
func SendVerificationCode(ctx context.Context, to string) (string, error) {
	code := generateVerificationCode()

	err := sendVerificationCode(to, code)
	if err != nil {
		return "", err
	}
	return code, nil
}

// sendVerificationCode 发送验证代码到指定的邮箱。
// 参数 to: 邮件接收人的邮箱地址。
// 参数 code: 需要发送的验证代码。
// 返回值 error: 发送过程中遇到的任何错误。
func sendVerificationCode(to string, code string) error {
	// 创建一个新的邮件实例
	conf := config.GetConfig().Email
	em := email.NewEmail()
	em.From = fmt.Sprintf("<%s> %s", conf.NickName, conf.From)
	em.To = []string{to}
	em.Subject = ""
	// 设置邮件的HTML内容
	em.HTML = []byte(`
		<h1>邮箱绑定验证码</h1>
		<p>您的邮箱验证码为: <strong>` + code + `</strong></p>
	`)
	// 发送邮件(这里使用QQ进行发送邮件验证码)
	err := em.Send(fmt.Sprintf("%s:%d", conf.Host, conf.Port), smtp.PlainAuth("", conf.From, conf.Secret, conf.Host))
	if err != nil {
		return err // 如果发送过程中有错误，返回错误信息
	}
	return nil // 邮件发送成功，返回nil
}

// 随机生成一个6位数的验证码。
func generateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	return code
}
