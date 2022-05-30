package openweather

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type geocoding struct {
	opt    *options
	sender *sender
}

func newGeocoding(client *http.Client, opt *options) *geocoding {
	return &geocoding{
		opt:    opt,
		sender: newSender(client),
	}
}

func (g *geocoding) GeocodingByCoordinates(ctx context.Context, model *ListGeocodingData, lat, lon float64, limit uint64) error {
	if err := ValidCoordinates(lat, lon); err != nil {
		return err
	}
	values := url.Values{}
	values.Add("lat", fmt.Sprintf("%f", lat))
	values.Add("lon", fmt.Sprintf("%f", lon))
	values.Add("limit", fmt.Sprintf("%d", limit))
	if err := g.sender.do(ctx, http.MethodGet, g.buildPath("reverse?", values), model, nil); err != nil {
		return err
	}
	return nil
}

func (g *geocoding) GeocodingByCityName(ctx context.Context, model *ListGeocodingData, q string, limit uint64) error {
	values := url.Values{}
	values.Add("q", q)
	values.Add("limit", fmt.Sprintf("%d", limit))
	if err := g.sender.do(ctx, http.MethodGet, g.buildPath("direct?", values), model, nil); err != nil {
		return err
	}
	return nil
}

func (g *geocoding) GeocodingByZip(ctx context.Context, model *ZipGeocodingData, zip string) error {
	values := url.Values{}
	values.Add("zip", zip)
	if err := g.sender.do(ctx, http.MethodGet, g.buildPath("zip?", values), model, nil); err != nil {
		return err
	}
	return nil
}

func (g *geocoding) buildPath(path string, values url.Values) string {
	values.Add("appid", g.opt.key)
	return geoUrl + path + values.Encode()
}
