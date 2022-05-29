package openweather

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type forecast struct {
	opt    *options
	sender *sender
}

func newForecast(client *http.Client, opt *options) *forecast {
	return &forecast{
		opt:    opt,
		sender: newSender(client),
	}
}

func (f *forecast) ForecastByCoordinates(ctx context.Context, model *ForecastData, lat, lon float64, cnt int) error {
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

func (f *forecast) ForecastByCityName(ctx context.Context, model *ForecastData, q string, cnt int) error {
	values := url.Values{}
	values.Add("q", q)
	values.Add("cnt", fmt.Sprintf("%d", cnt))
	if err := f.sender.do(ctx, http.MethodGet, f.buildPath(values), model, nil); err != nil {
		return err
	}
	return nil
}

func (f *forecast) ForecastByCityId(ctx context.Context, model *ForecastData, id int64, cnt int) error {
	values := url.Values{}
	values.Add("id", fmt.Sprintf("%d", id))
	values.Add("cnt", fmt.Sprintf("%d", cnt))
	if err := f.sender.do(ctx, http.MethodGet, f.buildPath(values), model, nil); err != nil {
		return err
	}
	return nil
}

func (f *forecast) ForecastByZip(ctx context.Context, model *ForecastData, zip string, cnt int) error {
	values := url.Values{}
	values.Add("zip", zip)
	values.Add("cnt", fmt.Sprintf("%d", cnt))
	if err := f.sender.do(ctx, http.MethodGet, f.buildPath(values), model, nil); err != nil {
		return err
	}
	return nil
}

func (f *forecast) buildPath(values url.Values) string {
	for k, v := range f.opt.getMap() {
		values.Add(k, v)
	}
	return baseUrl + "forecast?" + values.Encode()
}
