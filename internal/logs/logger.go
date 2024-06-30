package logs

import (
	"context"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 定义一个全局的logger变量
var logger *zap.Logger

// 初始化函数
func init() {
	// 获取日志写入器
	writeSyncer := getLogWriter(nil)
	// 获取编码器
	encoder := getEncoder()
	// 创建一个core
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// 创建logger
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// 获取编码器
func getEncoder() zapcore.Encoder {
	// 创建编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 设置级别格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 返回编码器
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 设置日志写入器
func SetLogerWriter(writer io.Writer) {
	// 获取日志写入器
	writeSyncer := getLogWriter(writer)
	// 获取编码器
	encoder := getEncoder()
	// 创建core
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// 创建logger
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// 获取日志写入器
func getLogWriter(writer io.Writer) zapcore.WriteSyncer {
	// 如果writer为空，则返回标准输出
	if writer == nil {
		return zapcore.AddSync(os.Stdout)
	}
	// 否则返回传入的writer
	return zapcore.AddSync(writer)
}

// 解析字段
func parseFields(ctx context.Context, fields ...zap.Field) []zap.Field {
	// 获取traceid
	traceID := ctx.Value("traceid")
	// 如果traceid不为空，则将其添加到fields中
	if traceID != nil {
		fields = append(fields, zap.String("traceid", traceID.(string)))
	}
	// 返回fields
	return fields
}

// Debug函数
func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	// 调用logger的Debug方法
	logger.Debug(msg, parseFields(ctx, fields...)...)
}

// Info函数
func Info(ctx context.Context, msg string, fields ...zap.Field) {
	// 调用logger的Info方法
	logger.Info(msg, parseFields(ctx, fields...)...)
}

// Warn函数
func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	// 调用logger的Warn方法
	logger.Warn(msg, parseFields(ctx, fields...)...)
}

// Error函数
func Error(ctx context.Context, msg string, fields ...zap.Field) {
	// 调用logger的Error方法
	logger.Error(msg, parseFields(ctx, fields...)...)
}

// Panic函数
func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	// 调用logger的Panic方法
	logger.Panic(msg, parseFields(ctx, fields...)...)
}

// Fatal函数
func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	// 调用logger的Fatal方法
	logger.Fatal(msg, parseFields(ctx, fields...)...)
}
