package handlers

import (
	"net/url"
	"testing"
)

func TestParseCoordinates(t *testing.T) {
	data := []struct {
		name   string
		query  url.Values
		lat    float64
		lon    float64
		errMsg string
	}{
		{
			"latitude is empty",
			map[string][]string{
				"lat": {""},
				"lon": {""},
			},
			0,
			0,
			"empty latitude",
		},
		{
			"longitude is empty",
			map[string][]string{
				"lat": {"10"},
				"lon": {""},
			},
			0,
			0,
			"empty longitude",
		},
		{
			"latitude is invalid",
			map[string][]string{
				"lat": {"ttt"},
				"lon": {"10"},
			},
			0,
			0,
			"invalid latitude",
		},
		{
			"longitude is invalid",
			map[string][]string{
				"lat": {"10"},
				"lon": {"ttt"},
			},
			0,
			0,
			"invalid longitude",
		},
		{
			"successful test",
			map[string][]string{
				"lat": {"10"},
				"lon": {"10"},
			},
			10,
			10,
			"",
		},
	}

	p := parser{}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			lat, lon, err := p.parseCoordinates(d.query)

			if lat != d.lat || lon != d.lon {
				t.Errorf("Expected coordinates lat: `%f` lon: `%f`, got lat: `%f` lon: `%f`",
					d.lat, d.lon, lat, lon)
			}

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}

			if errMsg != d.errMsg {
				t.Errorf("Expected error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}

func TestParseCityName(t *testing.T) {
	data := []struct {
		name   string
		query  url.Values
		q      string
		errMsg string
	}{
		{
			"city name is empty",
			map[string][]string{
				"q": {""},
			},
			"",
			"empty q",
		},
		{
			"successful test",
			map[string][]string{
				"q": {"London"},
			},
			"London",
			"",
		},
	}

	p := parser{}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			q, err := p.parseCityName(d.query)

			if q != d.q {
				t.Errorf("Expected city name `%s`, got `%s`",
					d.q, q)
			}

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}

			if errMsg != d.errMsg {
				t.Errorf("Expected error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}

func TestParseCityID(t *testing.T) {
	data := []struct {
		name   string
		query  url.Values
		id     uint64
		errMsg string
	}{
		{
			"city id is empty",
			map[string][]string{
				"id": {""},
			},
			0,
			"empty id",
		},
		{
			"city id is invalid",
			map[string][]string{
				"id": {"ttt"},
			},
			0,
			"invalid id",
		},
		{
			"successful test",
			map[string][]string{
				"id": {"10"},
			},
			10,
			"",
		},
	}

	p := parser{}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			id, err := p.parseCityID(d.query)

			if id != d.id {
				t.Errorf("Expected id `%d`, got `%d`",
					d.id, id)
			}

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}

			if errMsg != d.errMsg {
				t.Errorf("Expected error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}

func TestZipCode(t *testing.T) {
	data := []struct {
		name   string
		query  url.Values
		errMsg string
	}{
		{
			"zip is empty",
			map[string][]string{
				"zip": {""},
			},
			"empty zip",
		},
		{
			"successful test",
			map[string][]string{
				"zip": {"10"},
			},
			"",
		},
	}

	p := parser{}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			_, err := p.parseZipCode(d.query)

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}

			if errMsg != d.errMsg {
				t.Errorf("Expected error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
