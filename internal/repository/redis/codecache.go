package redis

import (
	"context"
	"fmt"
	"time"
)

type CodeCache struct {
}

func NewCodeCache() *CodeCache {
	return &CodeCache{}
}

func (u *CodeCache) SetPhoneCode(feild string, codekey string, code string, expire int64) error {
	key := fmt.Sprintf("%s:%s", feild, codekey)
	return _cache.Set(context.Background(), key, code, time.Duration(expire)*time.Minute).Err()
}

func (u *CodeCache) GetPhoneCode(feild string, codekey string) (string, error) {
	key := fmt.Sprintf("%s:%s", feild, codekey)

	return _cache.Get(context.Background(), key).Result()
}

func (u *CodeCache) SetEmailCode(feild string, codekey string, code string, expire int64) error {
	key := fmt.Sprintf("%s:%s", feild, codekey)
	return _cache.Set(context.Background(), key, code, time.Duration(expire)*time.Minute).Err()
}

const filterName = "myfilter"

func (u *CodeCache) GetEmailCode(feild string, codekey string) (string, error) {
	key := fmt.Sprintf("%s:%s", feild, codekey)
	_cache.Do(context.Background(), "BF.ADD", filterName, "element1")
	return _cache.Get(context.Background(), key).Result()
}
