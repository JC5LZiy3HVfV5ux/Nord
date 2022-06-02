package repositories

import (
	"context"
	"time"

	"github.com/JC5LZiy3HVfV5ux/openweather-cache-server/openweather"
	"github.com/go-redis/redis/v8"
)

type CurrentWeather interface {
	CurrentGet(ctx context.Context, model *openweather.CurrentWeatherData, key string) error
	CurrentPut(ctx context.Context, model *openweather.CurrentWeatherData, key string) error
}

type Forecast interface {
	ForecastGet(ctx context.Context, model *openweather.ForecastData, key string) error
	ForecastPut(ctx context.Context, model *openweather.ForecastData, key string) error
}

type Repository struct {
	CurrentWeather
	Forecast
}

func NewRepository(db *redis.Client, expiration time.Duration) *Repository {
	return &Repository{
		CurrentWeather: NewCurrentWeatherRepository(db, expiration),
		Forecast:       NewForecastRepository(db, expiration),
	}
}
