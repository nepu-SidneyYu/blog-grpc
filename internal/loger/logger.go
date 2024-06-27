package loger

import (
	"io"
	"log"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

var (
	_writer io.Writer
)

// SetWriter set writer
func Init() {
	_logwriter, err := rotatelogs.New(
		"./logs/blog-grpc_%Y%m%d%H%M.log",
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 最长保存30天
		rotatelogs.WithRotationTime(time.Hour*24), // 24小时切割一次
	)
	if err != nil {
		panic(err)
	}
	_writer = io.MultiWriter(_logwriter, os.Stdout)
	log.SetOutput(_writer)
}

func Writer() io.Writer {
	return _writer
}
