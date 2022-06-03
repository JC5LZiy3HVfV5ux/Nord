package repositories

import (
	"context"
	"time"

	"github.com/JC5LZiy3HVfV5ux/nord/pkg/openweather"
	"github.com/go-redis/redis/v8"
)

type ForecastRepository struct {
	redis      *redisClient
	expiration time.Duration
}

func NewForecastRepository(db *redis.Client, expiration time.Duration) *ForecastRepository {
	return &ForecastRepository{
		redis:      newRedisClient(db),
		expiration: expiration,
	}
}

func (f *ForecastRepository) ForecastGet(ctx context.Context, model *openweather.ForecastData, key string) error {
	if err := f.redis.get(ctx, model, key); err != nil {
		return err
	}

	return nil
}

func (f *ForecastRepository) ForecastPut(ctx context.Context, model *openweather.ForecastData, key string) error {
	if err := f.redis.put(ctx, model, key, f.expiration); err != nil {
		return err
	}

	return nil
}
