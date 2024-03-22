package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var _ RedisRepository = (*RedisRepo)(nil)

type RedisRepository interface {
	Get(ctx context.Context, key string) (uint64, error)
	Increment(ctx context.Context, key string) (value uint64, err error)
	Delete(ctx context.Context, key string) error
}

type RedisRepo struct {
	*redis.Client
}

func New(client *redis.Client) *RedisRepo {
	return &RedisRepo{
		Client: client,
	}
}
