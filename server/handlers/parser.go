package handlers

import (
	"errors"
	"net/url"
	"strconv"
)

type parser struct{}

func (p parser) parseCoordinates(query url.Values) (lat, lon float64, err error) {
	if query.Get("lat") == "" {
		return 0, 0, errors.New("empty latitude")
	}

	if query.Get("lon") == "" {
		return 0, 0, errors.New("empty longitude")
	}

	if lat, err = strconv.ParseFloat(query.Get("lat"), 64); err != nil {
		return 0, 0, errors.New("invalid latitude")
	}

	if lon, err = strconv.ParseFloat(query.Get("lon"), 64); err != nil {
		return 0, 0, errors.New("invalid longitude")
	}

	return
}

func (p parser) parseCityName(query url.Values) (q string, err error) {
	if q = query.Get("q"); q == "" {
		return "", errors.New("empty q")
	}

	return q, nil
}

func (p parser) parseCityID(query url.Values) (id uint64, err error) {
	if query.Get("id") == "" {
		return 0, errors.New("empty id")
	}

	if id, err = strconv.ParseUint(query.Get("id"), 10, 64); err != nil {
		return 0, errors.New("invalid id")
	}

	return id, nil
}

func (p parser) parseZipCode(query url.Values) (zip string, err error) {
	if zip = query.Get("zip"); zip == "" {
		return "", errors.New("empty zip")
	}

	return zip, nil
}

func (p parser) parseCnt(query url.Values) (cnt uint64, err error) {
	if query.Get("cnt") == "" {
		return 40, nil
	}

	if cnt, err = strconv.ParseUint(query.Get("cnt"), 10, 64); err != nil {
		return 0, errors.New("invalid cnt")
	}

	return cnt, nil
}
