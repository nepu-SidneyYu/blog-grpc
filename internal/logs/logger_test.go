package logs

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	ctx := context.Background()
	Info(ctx, "test info log")
	Info(ctx, "test info string log", String("key", "value"))

	Debug(ctx, "test debug int log",
		Int("int", 1),
		Int8("string", 2),
		Int16("int16", 3),
		Int32("int32", 4),
		Int64("int64", 5),
		Uint("uint", 6),
		Uint8("uint8", 7),
		Uint16("uint16", 8),
		Uint32("uint32", 9),
		Uint64("uint64", 10),
	)

	Warn(ctx, "test warn float log",
		Float32("float32", 1.1),
		Float64("float64", 2.2),
	)

	Error(ctx, "test error bool & error log",
		Bool("bool", true),
		Err(errors.New("test error")),
	)

	Debug(ctx, "test debug any log",
		Any("any", "any"),
	)

	Debug(ctx, "test debug binary log",
		Binary("binary", []byte("binary")),
		ByteString("byteString", []byte("byteString")),
	)

	Debug(ctx, "test debug time log",
		Time("time", time.Now()),
		Duration("duration", time.Second),
	)

	Info(ctx, "test info stack log", Stack("stack"))

	// Panic(ctx, "test panic any log",
	// 	Any("any", "any"),
	// )

	// Fatal(ctx, "test fatal binary log",
	// 	Binary("binary", []byte("binary")),
	// 	ByteString("byteString", []byte("byteString")),
	// )
}

func TestLogWithTrace(t *testing.T) {
	ctx := context.WithValue(context.Background(), "traceid", "123456")
	Info(ctx, "test info log")
	Info(ctx, "test info string log", String("key", "value"))

	Debug(ctx, "test debug int log",
		Int("int", 1),
		Int8("string", 2),
		Int16("int16", 3),
		Int32("int32", 4),
		Int64("int64", 5),
		Uint("uint", 6),
		Uint8("uint8", 7),
		Uint16("uint16", 8),
		Uint32("uint32", 9),
		Uint64("uint64", 10),
	)

	Warn(ctx, "test warn float log",
		Float32("float32", 1.1),
		Float64("float64", 2.2),
	)

	Error(ctx, "test error bool & error log",
		Bool("bool", true),
		Err(errors.New("test error")),
	)

	Debug(ctx, "test debug any log",
		Any("any", "any"),
	)

	Debug(ctx, "test debug binary log",
		Binary("binary", []byte("binary")),
		ByteString("byteString", []byte("byteString")),
	)

	Debug(ctx, "test debug time log",
		Time("time", time.Now()),
		Duration("duration", time.Second),
	)

	Info(ctx, "test info stack log", Stack("stack"))

	// Panic(ctx, "test panic any log",
	// 	Any("any", "any"),
	// )

	// Fatal(ctx, "test fatal binary log",
	// 	Binary("binary", []byte("binary")),
	// 	ByteString("byteString", []byte("byteString")),
	// )
}
