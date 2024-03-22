package redis

import "context"

func (r *RedisRepo) Increment(ctx context.Context, key string) (value uint64, err error) {
	return r.Client.Incr(ctx, key).Uint64()
}
