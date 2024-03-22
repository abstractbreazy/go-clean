package limiter

import (
	"context"

	"go-clean/internal/config"
)

type redisRepository interface {
	Get(ctx context.Context, key string) (uint64, error)
	Increment(ctx context.Context, key string) (value uint64, err error)
	Delete(ctx context.Context, key string) error
}

type limiterRepository interface {
	GetSubnet(ip string) (string, error)
}

type Service struct {
	cfg               *config.Config
	redisRepository   redisRepository
	limiterRepository limiterRepository
}

func NewService(
	cfg *config.Config,
	redisRepository redisRepository,
	limiterRepository limiterRepository,
) *Service {
	return &Service{
		cfg:               cfg,
		redisRepository:   redisRepository,
		limiterRepository: limiterRepository,
	}
}
