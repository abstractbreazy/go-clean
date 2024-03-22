package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go-clean/internal/config"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func NewServer(cfg *config.HTTP, router http.Handler) *Server {
	server := &Server{}
	server.echo = echo.New()
	server.echo.Server = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	return server
}

func (s *Server) RegisterRoutes(r http.Handler) {
	s.echo.Server.Handler = r
}

func (s *Server) Run() {
	go s.run()
}

func (s *Server) run() {
	if err := s.echo.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func (s *Server) Close() (err error) {
	return s.echo.Close()
}
