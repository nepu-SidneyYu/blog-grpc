package config

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// 调用LoadConfig函数进行测试
	LoadConfig()
	fmt.Printf("%#v\n", GetConfig())
}
