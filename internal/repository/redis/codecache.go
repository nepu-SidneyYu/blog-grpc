package redis

import (
	"context"
	"time"
)

type CodeCache struct {
}

func NewCodeCache() *CodeCache {
	return &CodeCache{}
}

func (u *CodeCache) SetCode(codekey string, code string, expire int64) error {
	return _cache.Set(context.Background(), codekey, code, time.Duration(expire)*time.Minute).Err()
}

func (u *CodeCache) GetCode(codekey string) (string, error) {
	return _cache.Get(context.Background(), codekey).Result()
}
