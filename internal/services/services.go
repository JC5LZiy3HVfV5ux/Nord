package services

import (
	"context"
	"errors"

	"github.com/JC5LZiy3HVfV5ux/nord/internal/repositories"
	"github.com/JC5LZiy3HVfV5ux/nord/pkg/openweather"
)

var (
	errResponseService = errors.New("internal server error")
	errOpenweather     *openweather.HttpError
)

type CurrentWeather interface {
	CurrentByCoordinates(ctx context.Context, lat, lon float64) (*openweather.CurrentWeatherData, error)
	CurrentByCityName(ctx context.Context, q string) (*openweather.CurrentWeatherData, error)
	CurrentByCityId(ctx context.Context, id uint64) (*openweather.CurrentWeatherData, error)
	CurrentByZip(ctx context.Context, zip string) (*openweather.CurrentWeatherData, error)
}

type Forecast interface {
	ForecastByCoordinates(ctx context.Context, lat, lon float64, cnt uint64) (*openweather.ForecastData, error)
	ForecastByCityName(ctx context.Context, q string, cnt uint64) (*openweather.ForecastData, error)
	ForecastByCityId(ctx context.Context, id uint64, cnt uint64) (*openweather.ForecastData, error)
	ForecastByZip(ctx context.Context, zip string, cnt uint64) (*openweather.ForecastData, error)
}

type Services struct {
	CurrentWeather
	Forecast
}

func NewServices(openweather *openweather.Client, repository *repositories.Repository) *Services {
	return &Services{
		CurrentWeather: NewCurrentWeatherService(openweather.CurrentWeather(), repository.CurrentWeather),
		Forecast:       NewForecastService(openweather.Forecast(), repository.Forecast),
	}
}
