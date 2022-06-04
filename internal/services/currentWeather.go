package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/JC5LZiy3HVfV5ux/nord/internal/repositories"
	"github.com/JC5LZiy3HVfV5ux/nord/pkg/openweather"
)

type CurrentWeatherService struct {
	currentWeather *openweather.CurrentWeather
	repository     repositories.CurrentWeather
}

func NewCurrentWeatherService(currentWeather *openweather.CurrentWeather, repository repositories.CurrentWeather) *CurrentWeatherService {
	return &CurrentWeatherService{
		currentWeather: currentWeather,
		repository:     repository,
	}
}

func (c *CurrentWeatherService) CurrentByCoordinates(ctx context.Context, lat, lon float64) (*openweather.CurrentWeatherData, error) {
	model := &openweather.CurrentWeatherData{}

	if err := c.repository.CurrentGet(ctx, model, fmt.Sprintf("%f,%f", lat, lon)); err != nil {
		log.Println(err)

		if err := c.currentWeather.CurrentByCoordinates(ctx, model, lat, lon); err != nil {
			log.Println(err)

			if errors.As(err, &errOpenweather) {
				return nil, err
			}

			return nil, errResponseService
		}

		if err := c.repository.CurrentPut(ctx, model, fmt.Sprintf("%f,%f", lat, lon)); err != nil {
			log.Println(err)
		}
	}

	return model, nil
}

func (c *CurrentWeatherService) CurrentByCityName(ctx context.Context, q string) (*openweather.CurrentWeatherData, error) {
	model := &openweather.CurrentWeatherData{}

	if err := c.repository.CurrentGet(ctx, model, q); err != nil {
		log.Println(err)

		if err := c.currentWeather.CurrentByCityName(ctx, model, q); err != nil {
			log.Println(err)

			if errors.As(err, &errOpenweather) {
				return nil, err
			}

			return nil, errResponseService
		}

		if err := c.repository.CurrentPut(ctx, model, q); err != nil {
			log.Println(err)
		}
	}

	return model, nil
}

func (c *CurrentWeatherService) CurrentByCityId(ctx context.Context, id uint64) (*openweather.CurrentWeatherData, error) {
	model := &openweather.CurrentWeatherData{}

	if err := c.repository.CurrentGet(ctx, model, fmt.Sprintf("%d", id)); err != nil {
		log.Println(err)

		if err := c.currentWeather.CurrentByCityId(ctx, model, id); err != nil {
			log.Println(err)

			if errors.As(err, &errOpenweather) {
				return nil, err
			}

			return nil, errResponseService
		}

		if err := c.repository.CurrentPut(ctx, model, fmt.Sprintf("%d", id)); err != nil {
			log.Println(err)
		}
	}

	return model, nil
}

func (c *CurrentWeatherService) CurrentByZip(ctx context.Context, zip string) (*openweather.CurrentWeatherData, error) {
	model := &openweather.CurrentWeatherData{}

	if err := c.repository.CurrentGet(ctx, model, zip); err != nil {
		log.Println(err)

		if err := c.currentWeather.CurrentByZip(ctx, model, zip); err != nil {
			log.Println(err)

			if errors.As(err, &errOpenweather) {
				return nil, err
			}

			return nil, errResponseService
		}

		if err := c.repository.CurrentPut(ctx, model, zip); err != nil {
			log.Println(err)
		}
	}

	return model, nil
}
