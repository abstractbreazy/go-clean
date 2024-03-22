package app

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/go-redis/redis/v8"

	"go-clean/internal/config"
)

// The Redis represents Redis connection.
type Redis struct {
	*redis.Client
}

// New Redis this given configuration.
func (a *App) newRedis(cfg config.Redis) (*Redis, error) {
	var (
		rd  = new(Redis)
		err error
		u   *url.URL
	)

	if u, err = url.Parse(cfg.URL); err != nil {
		return nil, fmt.Errorf("can't parse Redis URL: %w", err)
	}

	switch u.Scheme {
	case "redis", "rediss":
		var optsRedis *redis.Options
		optsRedis, err = redis.ParseURL(cfg.URL)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Redis options: %s", err)
		}

		rd.Client = redis.NewClient(optsRedis)
		err = rd.Connect(context.Background())
	default:
		err = fmt.Errorf("not supported database with schema '%s'", u.Scheme)
	}

	return rd, err
}

// Connect to Redis server and ping.
func (r *Redis) Connect(ctx context.Context) error {
	if r.Client == nil {
		return errors.New("can't connect to redis")
	}
	err := r.Ping(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Pings redis server.
func (r *Redis) Ping(ctx context.Context) error {
	if r.Client == nil {
		return errors.New("client is not defined")
	}
	_, err := r.Client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to ping Redis client: %w", err)
	}
	return nil
}

// Close the entire redis server.
func (r *Redis) Close() error {
	return r.Client.Close()
}
