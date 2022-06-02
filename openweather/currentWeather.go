package openweather

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type CurrentWeather struct {
	opt    *options
	sender *sender
}

func newCurrentWeather(client *http.Client, opt *options) *CurrentWeather {
	return &CurrentWeather{
		opt:    opt,
		sender: newSender(client),
	}
}

func (c *CurrentWeather) CurrentByCoordinates(ctx context.Context, model *CurrentWeatherData, lat, lon float64) error {
	if err := ValidCoordinates(lat, lon); err != nil {
		return err
	}
	values := url.Values{}
	values.Add("lat", fmt.Sprintf("%f", lat))
	values.Add("lon", fmt.Sprintf("%f", lon))
	if err := c.sender.do(ctx, http.MethodGet, c.buildPath(values), model, nil); err != nil {
		return err
	}
	return nil
}

func (c *CurrentWeather) CurrentByCityName(ctx context.Context, model *CurrentWeatherData, q string) error {
	values := url.Values{}
	values.Add("q", q)
	if err := c.sender.do(ctx, http.MethodGet, c.buildPath(values), model, nil); err != nil {
		return err
	}
	return nil
}

func (c *CurrentWeather) CurrentByCityId(ctx context.Context, model *CurrentWeatherData, id uint64) error {
	values := url.Values{}
	values.Add("id", fmt.Sprintf("%d", id))
	if err := c.sender.do(ctx, http.MethodGet, c.buildPath(values), model, nil); err != nil {
		return err
	}
	return nil
}

func (c *CurrentWeather) CurrentByZip(ctx context.Context, model *CurrentWeatherData, zip string) error {
	values := url.Values{}
	values.Add("zip", zip)
	if err := c.sender.do(ctx, http.MethodGet, c.buildPath(values), model, nil); err != nil {
		return err
	}
	return nil
}

func (c *CurrentWeather) buildPath(values url.Values) string {
	for k, v := range c.opt.getMap() {
		values.Add(k, v)
	}
	return baseUrl + "weather?" + values.Encode()
}
