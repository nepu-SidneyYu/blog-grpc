package logs

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var (
	logger *zap.Logger
)

func init() {
	writeSyncer := getLogWriter(nil)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func SetLogerWriter(writer io.Writer) {
	writeSyncer := getLogWriter(writer)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func getLogWriter(writer io.Writer) zapcore.WriteSyncer {
	if writer == nil {
		return zapcore.AddSync(os.Stdout)
	}
	return zapcore.AddSync(writer)

}

func parseFields(ctx context.Context, fields ...zap.Field) []zap.Field {
	traceID := ctx.Value("traceid")
	if traceID != nil {
		fields = append(fields, zap.String("traceid", traceID.(string)))
	}
	return fields
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Debug(msg, parseFields(ctx, fields...)...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Info(msg, parseFields(ctx, fields...)...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Warn(msg, parseFields(ctx, fields...)...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Error(msg, parseFields(ctx, fields...)...)
}

func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Panic(msg, parseFields(ctx, fields...)...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	logger.Fatal(msg, parseFields(ctx, fields...)...)
}
