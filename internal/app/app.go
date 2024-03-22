package app

import (
	"errors"
	"go-clean/internal/config"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

type App struct {
	c     *Container
	cfg   *config.Config
	redis *redis.Client
	http  *echo.Echo
}

var a *App

func NewApp() (*App, error) {
	var (
		cfg *config.Config
		err error
	)
	if cfg, err = config.New().Init(); err != nil {
		return nil, err
	}

	app := &App{
		cfg: cfg,
	}

	var redis *Redis
	if redis, err = app.newRedis(*cfg.Redis); err != nil {
		return nil, err
	}
	app.redis = redis.Client

	httpClient := app.newHttpClient()
	app.http = httpClient

	app.c = NewContainer(cfg, app.redis, app.http, cfg.Limiter.SubnetMask)

	return app, nil
}

func SetGlobalApp(app *App) {
	a = app
}

func GetGlobalApp() (*App, error) {
	if a == nil {
		return nil, errors.New("global app is not initialized")
	}

	return a, nil
}
