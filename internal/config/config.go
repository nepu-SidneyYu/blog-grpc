package config

import (
	"context"
	"fmt"
	"os"

	"github.com/nepu-SidneyYu/blog-grpc/internal/logs"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
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
}

var (
	_config Config
)

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

func LoadConfig() {
	data, err := os.ReadFile("conf.yaml")
	if err != nil {
		logs.Fatal(context.Background(), "read yaml failed", zap.String("error", err.Error()))
	}
	fmt.Printf("%s\n", string(data))
	err = yaml.Unmarshal(data, &_config)
	if err != nil {
		logs.Fatal(context.Background(), "unmarshal yaml failed", zap.String("error", err.Error()))
	}
}

func GetConfig() Config {
	return _config
}
