package sendcode

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/nepu-SidneyYu/blog-grpc/internal/config"
	"github.com/nepu-SidneyYu/blog-grpc/internal/logs"
	"go.uber.org/zap"
)

func GenerateRandomCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func SendPhoneCode(phone string) (string, error) {
	conf := config.GetConfig().Phone
	// 生成6位随机Code
	code := GenerateRandomCode(conf.CodeNum)
	// 通过accessKey Id和Secret连接服务
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", conf.AccessKeyId, conf.AccessKeySecret)
	if err != nil {
		return "", err
	}
	request := dysmsapi.CreateSendSmsRequest() //创建请求
	request.Scheme = conf.Scheme               //请求协议，可选：https，但会慢一点
	request.PhoneNumbers = phone               //接收短信的手机号码
	request.SignName = conf.SignName           //短信签名名称
	request.TemplateCode = conf.TemplateCode   //短信模板ID
	Param, err := json.Marshal(map[string]interface{}{
		"code": code, // 验证码参数
	})
	if err != nil {
		logs.Error(context.Background(), "Marshall SendSMS Param error:", zap.String("error", err.Error()))
		return "", err
	}
	request.TemplateParam = string(Param) //将短信模板参数传入短信模板
	_, err = client.SendSms(request)      //调用阿里云API发送信息
	if err != nil {
		logs.Error(context.Background(), "SendSMS error:", zap.String("error", err.Error()))
		return "", err
	}
	return code, nil
}
