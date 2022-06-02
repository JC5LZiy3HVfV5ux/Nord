package openweather

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type Forecast struct {
	opt    *options
	sender *sender
}

func newForecast(client *http.Client, opt *options) *Forecast {
	return &Forecast{
		opt:    opt,
		sender: newSender(client),
	}
}

func (f *Forecast) ForecastByCoordinates(ctx context.Context, model *ForecastData, lat, lon float64, cnt uint64) error {
	if err := ValidCoordinates(lat, lon); err != nil {
		return err
	}
	values := url.Values{}
	values.Add("lat", fmt.Sprintf("%f", lat))
	values.Add("lon", fmt.Sprintf("%f", lon))
	values.Add("cnt", fmt.Sprintf("%d", cnt))
	if err := f.sender.do(ctx, http.MethodGet, f.buildPath(values), model, nil); err != nil {
		return err
	}
	return nil
}

func (f *Forecast) ForecastByCityName(ctx context.Context, model *ForecastData, q string, cnt uint64) error {
	values := url.Values{}
	values.Add("q", q)
	values.Add("cnt", fmt.Sprintf("%d", cnt))
	if err := f.sender.do(ctx, http.MethodGet, f.buildPath(values), model, nil); err != nil {
		return err
	}
	return nil
}

func (f *Forecast) ForecastByCityId(ctx context.Context, model *ForecastData, id uint64, cnt uint64) error {
	values := url.Values{}
	values.Add("id", fmt.Sprintf("%d", id))
	values.Add("cnt", fmt.Sprintf("%d", cnt))
	if err := f.sender.do(ctx, http.MethodGet, f.buildPath(values), model, nil); err != nil {
		return err
	}
	return nil
}

func (f *Forecast) ForecastByZip(ctx context.Context, model *ForecastData, zip string, cnt uint64) error {
	values := url.Values{}
	values.Add("zip", zip)
	values.Add("cnt", fmt.Sprintf("%d", cnt))
	if err := f.sender.do(ctx, http.MethodGet, f.buildPath(values), model, nil); err != nil {
		return err
	}
	return nil
}

func (f *Forecast) buildPath(values url.Values) string {
	for k, v := range f.opt.getMap() {
		values.Add(k, v)
	}
	return baseUrl + "forecast?" + values.Encode()
}
