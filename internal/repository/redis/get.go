package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

func (r *RedisRepo) Get(ctx context.Context, key string) (uint64, error) {
	var (
		pipeline   = r.Client.Pipeline()
		value      = pipeline.Get(ctx, key)
		ttl        = pipeline.TTL(ctx, key)
		err        error
		actualTTL  time.Duration
		expiration time.Duration
		total      uint64
	)

	if _, err = pipeline.Exec(ctx); err != nil {
		if !errors.Is(err, redis.Nil) {
			return 0, err
		}
	}

	if actualTTL, err = ttl.Result(); err != nil || actualTTL == -1 || actualTTL == -2 {
		expiration = time.Second * 10
		if err = r.Client.Expire(ctx, key, expiration).Err(); err != nil {
			return 0, err
		}
	} else {
		expiration = actualTTL
	}

	total, err = value.Uint64()
	if err != nil && !errors.Is(err, redis.Nil) {
		return 0, err
	}

	return total, nil
}
