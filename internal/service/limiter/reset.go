package limiter

import (
	"context"

	"go-clean/internal/handler/http"
)

func (s *Service) Reset(ctx context.Context, ip string) (string, error) {
	var (
		subnet string
		err    error
	)

	if subnet, err = s.limiterRepository.GetSubnet(ip); err != nil {
		return "", http.ErrBadRequest
	}

	if err = s.redisRepository.Delete(ctx, subnet); err != nil {
		return "", http.ErrInternal
	}

	return subnet, nil
}
