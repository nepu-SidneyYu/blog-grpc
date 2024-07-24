package config

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/magiconair/properties"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/nepu-SidneyYu/blog-grpc/internal/logs"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// Config结构体，包含MySqlConfig字段
type Config struct {
	Port int `yaml:"port"`
	// MySql字段，指向MySqlConfig结构体
	MySql MySqlConfig `yaml:"mysql"`
	//jwt字段
	JWt   JWTConfig           `yaml:"jwt"`
	Redis RedisConfig         `yaml:"redis"`
	Email SendEmailCodeConfig `yaml:"email"`
	Phone SendPhoneCodeConfig `yaml:"phone"`
}
type NacosConfig struct {
	TimeoutMs            uint64 `yaml:"timeout_ms"`              // 超时时间
	NamespaceId          string `yaml:"namespace_id"`            // 命名空间
	CacheDir             string `yaml:"cache_dir"`               // 缓存目录
	NotLoadCacheAtStart  bool   `yaml:"not_load_cache_at_start"` // 是否在启动时加载缓存
	UpdateCacheWhenEmpty bool   `yaml:"update_cache_when_empty"` // 是否在缓存为空时更新缓存
	Username             string `yaml:"username"`                // 用户名
	Password             string `yaml:"password"`                // 密码
	LogDir               string `yaml:"log_dir"`                 // 日志目录
	LogLevel             string `yaml:"log_level"`               // 日志级别
	ContextPath          string `yaml:"context_path"`            // 上下文路径
	Scheme               string `yaml:"scheme"`                  // 协议
	IpAddr               string `yaml:"ip_addr"`                 // ip地址
	Port                 uint64 `yaml:"port"`                    // 端口
	DataType             string `yaml:"data_type"`               // 数据类型
	DataId               string `yaml:"data_id"`                 // 数据id
	Group                string `yaml:"group"`                   // 分组
}

var (
	_config Config
	_nacos  NacosConfig
)

type SendPhoneCodeConfig struct {
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	TemplateCode    string `yaml:"templateCode"`
	Expire          int32  `yaml:"expire"`
	CodeNum         int    `yaml:"codeNum"`
	SignName        string `yaml:"signName"`
	Scheme          string `yaml:"scheme"`
	Region          string `yaml:"region"`
}
type JWTConfig struct {
	Secret string `yaml:"secret"`
	Expire int32  `yaml:"expire"` // hour
	Issuer string `yaml:"issuer"`
}

type SendEmailCodeConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	From     string `yaml:"from"`
	Secret   string `yaml:"secret"`
	NickName string `yaml:"nickname"`
	Expire   int32  `yaml:"expire"` // minute
}

type MySqlConfig struct {
	// 数据库连接信息
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Charset  string `yaml:"charset"`
}
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func loadNacosConfig() {
	// TODO: 从nacos加载配置
	resp, err := http.Get("http://conf.zhihuishu.com/global/nacos.properties")
	if err != nil {
		logs.Fatal(context.Background(), "获取Nacos配置信息失败", zap.String("error", err.Error()))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Fatal(context.Background(), "读取Nacos配置信息失败", zap.String("error", err.Error()))
	}
	p, err := properties.Load(body, properties.UTF8)
	if err != nil {
		logs.Fatal(context.Background(), "解析Nacos配置信息失败", zap.String("error", err.Error()))
	}
	username := p.GetString("nacos.zhihuishu.server.username")
	password := p.GetString("nacos.zhihuishu.server.password")
	if username == "" || password == "" {
		logs.Fatal(context.Background(), "Nacos配置信息不完整", zap.String("error", "nacos.zhihuishu.server.username or nacos.zhihuishu.server.password is empty"))
	}
	_nacos = NacosConfig{Username: username, Password: password}
	logs.Info(context.Background(), "加载Nacos配置信息成功", zap.String("username", username), zap.String("password", password))
}

func LoadConfig() {
	loadNacosConfig()

	clientConfig := &constant.ClientConfig{
		NamespaceId:          "pract",
		NotLoadCacheAtStart:  true,
		LogDir:               "/tmp/nacos/log",
		CacheDir:             "/tmp/nacos/cache",
		TimeoutMs:            5000,
		Username:             _nacos.Username,
		Password:             _nacos.Password,
		UpdateCacheWhenEmpty: true,
	}
	serverConfigs := []constant.ServerConfig{
		{
			//ContextPath: "/nacos",
			IpAddr: "nacos-i.zhihuishu.com",
			Port:   8848,
			Scheme: "http",
		},
	}
	configClient, err := clients.NewConfigClient(
		vo.ConfigParam{
			ClientConfig:  clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		logs.Fatal(context.Background(), "获取项目启动配置信息失败", zap.String("error", err.Error()))
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "config.yaml",
		Group:  "config",
		Type:   "yaml",
	})
	if err != nil {
		logs.Fatal(context.Background(), "获取项目启动配置信息失败", zap.String("error", err.Error()))
	}
	logs.Info(context.Background(), "项目配置", zap.String("content", fmt.Sprintf("%s\n", content)))

	err = yaml.Unmarshal([]byte(content), &_config)
	if err != nil {
		logs.Fatal(context.Background(), "解析项目启动配置信息失败", zap.String("error", err.Error()))
	}
	logs.Info(context.Background(), "加载项目启动配置信息成功", zap.String("_config", fmt.Sprintf("%s\n", _config)))

	// data, err := os.ReadFile("conf.yaml")
	// if err != nil {
	// 	logs.Fatal(context.Background(), "read yaml failed", zap.String("error", err.Error()))
	// }
	// fmt.Printf("%s\n", string(data))
	// err = yaml.Unmarshal(data, &_config)
	// if err != nil {
	// 	logs.Fatal(context.Background(), "unmarshal yaml failed", zap.String("error", err.Error()))
	// }
}

func GetConfig() Config {
	return _config
}
