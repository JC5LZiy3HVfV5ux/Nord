package services

import (
	"context"
	"fmt"
	"log"

	"github.com/JC5LZiy3HVfV5ux/nord/internal/repositories"
	"github.com/JC5LZiy3HVfV5ux/nord/pkg/openweather"
)

type ForecastService struct {
	forecast   *openweather.Forecast
	repository repositories.Forecast
}

func NewForecastService(forecast *openweather.Forecast, repository repositories.Forecast) *ForecastService {
	return &ForecastService{
		forecast:   forecast,
		repository: repository,
	}
}

func (f *ForecastService) ForecastByCoordinates(ctx context.Context, lat, lon float64, cnt uint64) (*openweather.ForecastData, error) {
	if err := openweather.ValidCoordinates(lat, lon); err != nil {
		return nil, err
	}

	model := &openweather.ForecastData{}

	if err := f.repository.ForecastGet(ctx, model, fmt.Sprintf("%f,%f,%d", lat, lon, cnt)); err != nil {
		log.Println(err)

		if err := f.forecast.ForecastByCoordinates(ctx, model, lat, lon, cnt); err != nil {
			log.Println(err)

			if asErrOpenweather(err) {
				return nil, err
			}

			return nil, errResponseService
		}

		if err := f.repository.ForecastPut(ctx, model, fmt.Sprintf("%f,%f,%d", lat, lon, cnt)); err != nil {
			log.Println(err)
		}
	}

	return model, nil
}

func (f *ForecastService) ForecastByCityName(ctx context.Context, q string, cnt uint64) (*openweather.ForecastData, error) {
	model := &openweather.ForecastData{}

	if err := f.repository.ForecastGet(ctx, model, fmt.Sprintf("%s,%d", q, cnt)); err != nil {
		log.Println(err)

		if err := f.forecast.ForecastByCityName(ctx, model, q, cnt); err != nil {
			log.Println(err)

			if asErrOpenweather(err) {
				return nil, err
			}

			return nil, errResponseService
		}

		if err := f.repository.ForecastPut(ctx, model, fmt.Sprintf("%s,%d", q, cnt)); err != nil {
			log.Println(err)
		}
	}

	return model, nil
}

func (f *ForecastService) ForecastByCityId(ctx context.Context, id uint64, cnt uint64) (*openweather.ForecastData, error) {
	model := &openweather.ForecastData{}

	if err := f.repository.ForecastGet(ctx, model, fmt.Sprintf("%d,%d", id, cnt)); err != nil {
		log.Println(err)

		if err := f.forecast.ForecastByCityId(ctx, model, id, cnt); err != nil {
			log.Println(err)

			if asErrOpenweather(err) {
				return nil, err
			}

			return nil, errResponseService
		}

		if err := f.repository.ForecastPut(ctx, model, fmt.Sprintf("%d,%d", id, cnt)); err != nil {
			log.Println(err)
		}
	}

	return model, nil
}

func (f *ForecastService) ForecastByZip(ctx context.Context, zip string, cnt uint64) (*openweather.ForecastData, error) {
	model := &openweather.ForecastData{}

	if err := f.repository.ForecastGet(ctx, model, fmt.Sprintf("%s,%d", zip, cnt)); err != nil {
		log.Println(err)

		if err := f.forecast.ForecastByZip(ctx, model, zip, cnt); err != nil {
			log.Println(err)

			if asErrOpenweather(err) {
				return nil, err
			}

			return nil, errResponseService
		}

		if err := f.repository.ForecastPut(ctx, model, fmt.Sprintf("%s,%d", zip, cnt)); err != nil {
			log.Println(err)
		}
	}

	return model, nil
}
