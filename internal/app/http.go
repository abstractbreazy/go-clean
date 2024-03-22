package app

import (
	"context"
	"log"
	"time"

	"go-clean/internal/handler/http"
	v1 "go-clean/internal/handler/http/api/v1"
	"go-clean/pkg/sigint"

	"github.com/labstack/echo/v4"
)

func (a *App) StartHTTP() {
	var (
		handler = v1.NewHandler(a.c.GetLimiterService())
		router  = http.NewRouter(a.c.getHttp()).
			WithHandler(handler)
	)

	srv := http.NewServer(a.cfg.HTTP, router.Echo)
	srv.RegisterRoutes(router)

	log.Printf("starting HTTP server at %s:%s", a.cfg.HTTP.Host, a.cfg.HTTP.Port)

	// Run server in goroutine.
	srv.Run()

	log.Println("service is started:", a.cfg.HTTP.Port)

	sigint.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("stopping service ...")

	// Shutdown Server.
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) newHttpClient() *echo.Echo {
	return echo.New()
}
