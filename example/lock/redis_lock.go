package lock

import (
	_ "embed"

	"github.com/redis/go-redis/v9"
)

type redisLock struct {
	opts *Options
	rd   redis.UniversalClient
}

//go:embed scripts/refresh.lua
var refreshScript string

//go:embed scripts/unlock.lua
var unlockScript string

func NewRedisLock(opts ...OptionFunc) DistributedLock {
	r := &redisLock{}

	//设置部分属性
	for _, o := range opts {
		o(r.opts)
	}

	return r
}

func (r *redisLock) TryLock(key string) bool {
	return false
}

func (r *redisLock) Lock(key string) {
}

func (r *redisLock) UnLock(key string) bool {
	return false
}
