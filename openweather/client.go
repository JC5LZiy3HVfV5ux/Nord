package openweather

import (
	"errors"
	"net/http"
)

const (
	baseUrl = "https://api.openweathermap.org/data/2.5/"
	geoUrl  = "http://api.openweathermap.org/geo/1.0/"
)

type client struct {
	opt            *options
	currentWeather *currentWeather
	forecast       *forecast
	geocoding      *geocoding
}

func NewClient(key string, customer ...*http.Client) (*client, error) {
	if key == "" {
		return nil, errors.New("api key is empty")
	}

	if len(customer) == 0 {
		customer = append(customer, http.DefaultClient)
	}

	opt := initOptions(key)

	return &client{
		opt:            opt,
		currentWeather: newCurrentWeather(customer[0], opt),
		forecast:       newForecast(customer[0], opt),
		geocoding:      newGeocoding(customer[0], opt),
	}, nil
}

func (c *client) SetLang(lang string) error {
	if !ValidLang(lang) {
		return errors.New("invalid lang code")
	}
	c.opt.lang = lang
	return nil
}

func (c *client) SetUnit(unit string) error {
	if !ValidUnit(unit) {
		return errors.New("invalid unit")
	}
	c.opt.unit = unit
	return nil
}

func (c *client) CurrentWeather() *currentWeather {
	return c.currentWeather
}

func (c *client) Forecast() *forecast {
	return c.forecast
}

func (c *client) Geocoding() *geocoding {
	return c.geocoding
}
