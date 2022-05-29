package handlers

import (
	"net/url"
	"testing"
)

func TestParseCoordinates(t *testing.T) {
	data := []struct {
		name   string
		query  url.Values
		errMsg string
	}{
		{
			"latitude is empty",
			map[string][]string{
				"lat": {""},
				"lon": {""},
			},
			"empty latitude",
		},
		{
			"longitude is empty",
			map[string][]string{
				"lat": {"10"},
				"lon": {""},
			},
			"empty longitude",
		},
		{
			"latitude is invalid",
			map[string][]string{
				"lat": {"ttt"},
				"lon": {"10"},
			},
			"invalid latitude",
		},
		{
			"longitude is invalid",
			map[string][]string{
				"lat": {"10"},
				"lon": {"ttt"},
			},
			"invalid longitude",
		},
		{
			"successful test",
			map[string][]string{
				"lat": {"10"},
				"lon": {"10"},
			},
			"",
		},
	}

	p := parser{}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			_, _, err := p.parseCoordinates(d.query)

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
		errMsg string
	}{
		{
			"city name is empty",
			map[string][]string{
				"q": {""},
			},
			"empty q",
		},
		{
			"successful test",
			map[string][]string{
				"q": {"London"},
			},
			"",
		},
	}

	p := parser{}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			_, err := p.parseCityName(d.query)

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
		errMsg string
	}{
		{
			"city id is empty",
			map[string][]string{
				"id": {""},
			},
			"empty id",
		},
		{
			"city id is invalid",
			map[string][]string{
				"id": {"ttt"},
			},
			"invalid id",
		},
		{
			"successful test",
			map[string][]string{
				"id": {"10"},
			},
			"",
		},
	}

	p := parser{}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			_, err := p.parseCityID(d.query)

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
