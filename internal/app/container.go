package app

import (
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"

	"go-clean/internal/config"
	limiterRepo "go-clean/internal/repository/limiter"
	redisRepo "go-clean/internal/repository/redis"
	"go-clean/internal/service/limiter"
)

type Container struct {
	cfg     *config.Config
	redis   *redis.Client
	http    *echo.Echo
	netMask string
}

func NewContainer(
	cfg *config.Config,
	redis *redis.Client,
	http *echo.Echo,
	netMask string,
) *Container {
	return &Container{
		cfg:     cfg,
		redis:   redis,
		http:    http,
		netMask: netMask,
	}
}

func (c *Container) GetLimiterService() *limiter.Service {
	return limiter.NewService(
		c.getConfig(),
		c.getRedisRepository(),
		c.getLimiter(),
	)
}

func (c *Container) getConfig() *config.Config {
	return c.cfg
}

func (c *Container) getRedis() *redis.Client {
	return c.redis
}

func (c *Container) getNetMask() string {
	return c.netMask
}

func (c *Container) getHttp() *echo.Echo {
	return c.http
}

func (c *Container) getRedisRepository() *redisRepo.RedisRepo {
	return redisRepo.New(c.getRedis())
}

func (c *Container) getLimiter() *limiterRepo.Limiter {
	return limiterRepo.NewLimiter(c.getNetMask())
}
