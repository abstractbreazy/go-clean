package v1

import "context"

type Service interface {
	Ping(ctx context.Context, ip string) error
	Reset(ctx context.Context, ip string) (string, error)
}

type Handler struct {
	uc Service
}

func NewHandler(uc Service) *Handler {
	return &Handler{uc: uc}
}
