package redis

import (
	"context"
	"time"
)

func SetCode(codekey string, code string, expire int64) error {
	return _cache.Set(context.Background(), codekey, code, time.Duration(expire)*time.Minute).Err()
}
