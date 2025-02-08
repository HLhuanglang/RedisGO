package lock

import "time"

type Options struct {
	ExpireTime time.Duration //锁过期时间
}

type OptionFunc func(*Options)

func WithExpireTime(exptm int) OptionFunc {
	return func(o *Options) {
		o.ExpireTime = time.Duration(exptm)
	}
}
