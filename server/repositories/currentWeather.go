package repositories

import (
	"context"
	"time"

	"github.com/JC5LZiy3HVfV5ux/openweather-cache-server/openweather"
	"github.com/go-redis/redis/v8"
)

type CurrentWeatherRepository struct {
	redis      *redisClient
	expiration time.Duration
}

func NewCurrentWeatherRepository(db *redis.Client, expiration time.Duration) *CurrentWeatherRepository {
	return &CurrentWeatherRepository{
		redis:      newRedisClient(db),
		expiration: expiration,
	}
}

func (c *CurrentWeatherRepository) CurrentGet(ctx context.Context, model *openweather.CurrentWeatherData, key string) error {
	if err := c.redis.get(ctx, model, key); err != nil {
		return err
	}

	return nil
}

func (c *CurrentWeatherRepository) CurrentPut(ctx context.Context, model *openweather.CurrentWeatherData, key string) error {
	if err := c.redis.put(ctx, model, key, c.expiration); err != nil {
		return err
	}

	return nil
}
