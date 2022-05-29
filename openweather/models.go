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
	Coord      coord     `json:"coord"`
	Weather    []weather `json:"weather"`
	Base       string    `json:"base"`
	Main       main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       wind      `json:"wind"`
	Clouds     clouds    `json:"clouds"`
	Dt         int       `json:"dt"`
	Sys        sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

type ForecastData struct {
	Cod     string         `json:"cod"`
	Message float64        `json:"message"`
	Cnt     int            `json:"cnt"`
	List    []forecastList `json:"list"`
	City    city           `json:"city"`
}

type forecastList struct {
	Dt   int `json:"dt"`
	Main struct {
		main
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
		TempKf    float64 `json:"temp_kf"`
	} `json:"main"`
	Weather    []weather `json:"weather"`
	Clouds     clouds    `json:"clouds"`
	Wind       wind      `json:"wind"`
	Visibility int       `json:"visibility"`
	Pop        float64   `json:"pop"`
	Sys        struct {
		Pod string `json:"pod"`
	} `json:"sys"`
	DtTxt string `json:"dt_txt"`
}

type city struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Coord      coord  `json:"coord"`
	Country    string `json:"country"`
	Population int    `json:"population"`
	Timezone   int    `json:"timezone"`
	Sunrise    int    `json:"sunrise"`
	Sunset     int    `json:"sunset"`
}

type coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

func (w *weather) UrlIconWeather() string {
	return fmt.Sprintf(iconUrl, w.Icon)
}

type main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
	Gust  float64 `json:"gust"`
}

type clouds struct {
	All int `json:"all"`
}

type sys struct {
	Type    int    `json:"type"`
	Id      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

type geocodingData struct {
	Name       string            `json:"name"`
	LocalNames map[string]string `json:"local_names"`
	Lat        float64           `json:"lat"`
	Lon        float64           `json:"lon"`
	Country    string            `json:"country"`
	State      string            `json:"state"`
}

type ListGeocodingData []geocodingData

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
type httpError struct {
	Cod     customInt `json:"cod"`
	Message string    `json:"message"`
}

func (e *httpError) Error() string {
	return fmt.Sprintf("%d", e.Cod) + "," + e.Message
}

func httpErrorFromJson(r io.Reader) error {
	httpError := &httpError{}
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
