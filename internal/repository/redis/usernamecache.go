package redis

import (
	"context"

	"github.com/nepu-SidneyYu/blog-grpc/internal/consts"
)

type UsernameCache struct {
}

func NewUserNameCache() *UsernameCache {
	SetKeyAndErrorRate()
	return &UsernameCache{}
}

func SetKeyAndErrorRate() error {
	cmd := _cache.Do(context.Background(), "BF.RESERVE", consts.UserNameFiled, consts.UserNameErrorRate, consts.UserNameCampacity)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (u *UsernameCache) SetUserName(name string) error {
	_, err := _cache.BFAdd(context.Background(), consts.UserNameFiled, name).Result()
	if err != nil {
		return err
	}
	return nil
}
func (u *UsernameCache) IsUserNameExist(name string) bool {
	return _cache.BFExists(context.Background(), consts.UserNameFiled, name).Val()
}
