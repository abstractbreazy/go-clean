package limiter

import (
	"context"

	"go-clean/internal/handler/http"
)

func (s *Service) Ping(ctx context.Context, ip string) error {
	var (
		subnet string
		err    error
		total  uint64
		count  uint64
	)

	if subnet, err = s.limiterRepository.GetSubnet(ip); err != nil {
		return http.ErrBadRequest
	}

	if total, err = s.redisRepository.Get(ctx, subnet); err != nil {
		return http.ErrInternal
	}

	if total > s.cfg.Limiter.Limit {
		return http.ErrManyRequests
	}

	if count, err = s.redisRepository.Increment(ctx, subnet); err != nil {
		return http.ErrInternal
	}

	if count > s.cfg.Limiter.Limit {
		return http.ErrManyRequests
	}

	return nil
}
