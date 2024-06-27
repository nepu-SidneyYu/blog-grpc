package logs

import (
	"time"

	"go.uber.org/zap"
)

func Int(key string, value int) zap.Field {
	return zap.Int(key, value)
}

func Int8(key string, value int8) zap.Field {
	return zap.Int8(key, value)
}

func Int16(key string, value int16) zap.Field {
	return zap.Int16(key, value)
}

func Int32(key string, value int32) zap.Field {
	return zap.Int32(key, value)
}

func Int64(key string, value int64) zap.Field {
	return zap.Int64(key, value)
}

func Uint(key string, value uint) zap.Field {
	return zap.Uint(key, value)
}

func Uint8(key string, value uint8) zap.Field {
	return zap.Uint8(key, value)
}

func Uint16(key string, value uint16) zap.Field {
	return zap.Uint16(key, value)
}

func Uint32(key string, value uint32) zap.Field {
	return zap.Uint32(key, value)
}

func Uint64(key string, value uint64) zap.Field {
	return zap.Uint64(key, value)
}

func Float32(key string, value float32) zap.Field {
	return zap.Float32(key, value)
}

func Float64(key string, value float64) zap.Field {
	return zap.Float64(key, value)
}

func String(key string, value string) zap.Field {
	return zap.String(key, value)
}

func Bool(key string, value bool) zap.Field {
	return zap.Bool(key, value)
}

func Any(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func Err(err error) zap.Field {
	return zap.Error(err)
}

func ByteString(key string, value []byte) zap.Field {
	return zap.ByteString(key, value)
}

func Binary(key string, value []byte) zap.Field {
	return zap.Binary(key, value)
}

func Time(key string, t time.Time) zap.Field {
	return zap.Time(key, t)
}

func Duration(key string, d time.Duration) zap.Field {
	return zap.Duration(key, d)
}

func Stack(key string) zap.Field {
	return zap.StackSkip(key, 1)
}
