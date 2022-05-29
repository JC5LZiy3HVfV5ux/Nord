package openweather

import (
	"testing"
)

func TestValidLongitude(t *testing.T) {
	data := []struct {
		name string
		lon  float64
		ok   bool
	}{
		{"longitude > 180", 181, false},
		{"longitude < -180", -181, false},
		{"longitude = 180", 180, true},
		{"longitude = -180", -180, true},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			if ok := ValidLongitude(d.lon); ok != d.ok {
				t.Errorf("Expected `%t`, got `%t`",
					d.ok, ok)
			}
		})
	}
}

func TestValidLatitude(t *testing.T) {
	data := []struct {
		name string
		lon  float64
		ok   bool
	}{
		{"longitude > 90", 91, false},
		{"longitude < -90", -91, false},
		{"longitude = 90", 90, true},
		{"longitude = -90", -90, true},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			if ok := ValidLatitude(d.lon); ok != d.ok {
				t.Errorf("Expected `%t`, got `%t`",
					d.ok, ok)
			}
		})
	}
}
