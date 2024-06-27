package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config结构体，包含MySqlConfig字段
type Config struct {
	// MySql字段，指向MySqlConfig结构体
	MySql *MySqlConfig
}

var (
	_config Config
)

type JWTConfig struct {
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

func LoadConfig() {
	data, err := os.ReadFile("./conf.yaml")
	if err != nil {
		fmt.Printf("读取yaml文件错误：%#v\n", err)
		logs.Fatal("读取yaml文件错误：%#v", err)
	}
	err = yaml.Unmarshal(data, &_config)
	if err != nil {
		fmt.Printf("umarshall数据失败：%#v\n", err)
	}
}

func GetConfig() Config {
	return _config
}
