package openweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

var iconUrl = "https://openweathermap.org/img/wn/%s.png"

var errUnknownValueCustomInt = errors.New("customInt: parsing unknown value")

type CurrentWeatherData struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int       `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

type ForecastData struct {
	Cod     string         `json:"cod"`
	Message float64        `json:"message"`
	Cnt     int            `json:"cnt"`
	List    []ForecastList `json:"list"`
	City    City           `json:"city"`
}

type ForecastList struct {
	Dt   int `json:"dt"`
	Main struct {
		Main
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
		TempKf    float64 `json:"temp_kf"`
	} `json:"main"`
	Weather    []Weather `json:"weather"`
	Clouds     Clouds    `json:"clouds"`
	Wind       Wind      `json:"wind"`
	Visibility int       `json:"visibility"`
	Pop        float64   `json:"pop"`
	Sys        struct {
		Pod string `json:"pod"`
	} `json:"sys"`
	DtTxt string `json:"dt_txt"`
}

type City struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Coord      Coord  `json:"coord"`
	Country    string `json:"country"`
	Population int    `json:"population"`
	Timezone   int    `json:"timezone"`
	Sunrise    int    `json:"sunrise"`
	Sunset     int    `json:"sunset"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

func (w *Weather) UrlIconWeather() string {
	return fmt.Sprintf(iconUrl, w.Icon)
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	Id      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

type GeocodingData struct {
	Name       string            `json:"name"`
	LocalNames map[string]string `json:"local_names"`
	Lat        float64           `json:"lat"`
	Lon        float64           `json:"lon"`
	Country    string            `json:"country"`
	State      string            `json:"state"`
}

type ListGeocodingData []GeocodingData

type ZipGeocodingData struct {
	Zip     string  `json:"zip"`
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
}

//
// status code int or string
// {"cod":401, "message": "Invalid API key. Please see http://openweathermap.org/faq#error401 for more info."}
// {"cod":"400","message":"wrong longitude"}
//
type HttpError struct {
	Cod     customInt `json:"cod"`
	Message string    `json:"message"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func httpErrorFromJson(r io.Reader) error {
	httpError := &HttpError{}
	if err := json.NewDecoder(r).Decode(&httpError); err != nil {
		return err
	}
	return httpError
}

type customInt int

//
// https://pkg.go.dev/encoding/json#Unmarshal
// float64 default for JSON numbers
//
func (ci *customInt) UnmarshalJSON(data []byte) error {
	var raw interface{}
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	switch v := raw.(type) {
	case float64:
		*ci = customInt(v)
		return nil
	case string:
		parsed, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*ci = customInt(parsed)
		return nil
	default:
		return errUnknownValueCustomInt
	}
}
