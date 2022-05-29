package handlers

import (
	"errors"
	"net/url"
	"strconv"
)

type parser struct{}

func (p parser) parseCoordinates(query url.Values) (lat, lon float64, err error) {
	if query.Get("lat") == "" {
		err = errors.New("empty latitude")
		return
	}

	if query.Get("lon") == "" {
		err = errors.New("empty longitude")
		return
	}

	if lat, err = strconv.ParseFloat(query.Get("lat"), 64); err != nil {
		err = errors.New("invalid latitude")
		return
	}

	if lon, err = strconv.ParseFloat(query.Get("lon"), 64); err != nil {
		err = errors.New("invalid longitude")
		return
	}

	return
}

func (p parser) parseCityName(query url.Values) (q string, err error) {
	if q = query.Get("q"); q == "" {
		err = errors.New("empty q")
		return
	}

	return
}

func (p parser) parseCityID(query url.Values) (id int64, err error) {
	if query.Get("id") == "" {
		err = errors.New("empty id")
		return
	}

	if id, err = strconv.ParseInt(query.Get("id"), 10, 64); err != nil {
		err = errors.New("invalid id")
		return
	}

	return
}

func (p parser) parseZipCode(query url.Values) (zip string, err error) {
	if query.Get("zip") == "" {
		err = errors.New("empty zip")
		return
	}
	return
}
