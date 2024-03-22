package redis

import "context"

func (r *RedisRepo) Delete(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}
