package openweather

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestCurrentWeatherDataModel(t *testing.T) {
	currentWeatherDataJson := `{
		"coord": {
		  "lon": 139,
		  "lat": 35
		},
		"weather": [
		  {
			"id": 800,
			"main": "Clear",
			"description": "clear sky",
			"icon": "01n"
		  }
		],
		"base": "stations",
		"main": {
		  "temp": 283.88,
		  "feels_like": 282.52,
		  "temp_min": 283.88,
		  "temp_max": 283.88,
		  "pressure": 1011,
		  "humidity": 58
		},
		"visibility": 10000,
		"wind": {
		  "speed": 2.76,
		  "deg": 58,
		  "gust": 5.14
		},
		"clouds": {
		  "all": 0
		},
		"dt": 1647357046,
		"sys": {
		  "type": 2,
		  "id": 2019346,
		  "country": "JP",
		  "sunrise": 1647377649,
		  "sunset": 1647420692
		},
		"timezone": 32400,
		"id": 1851632,
		"name": "Shuzenji",
		"cod": 200
	  }`

	model := CurrentWeatherData{}
	if err := json.NewDecoder(strings.NewReader(currentWeatherDataJson)).Decode(&model); err != nil {
		t.Error(err)
	}
}

func TestForecastDataModel(t *testing.T) {
	forecastDataJson := `{
		"cod": "200",
		"message": 0,
		"cnt": 40,
		"list": [
		  {
			"dt": 1647345600,
			"main": {
			  "temp": 286.66,
			  "feels_like": 285.22,
			  "temp_min": 286.32,
			  "temp_max": 286.66,
			  "pressure": 1007,
			  "sea_level": 1007,
			  "grnd_level": 982,
			  "humidity": 44,
			  "temp_kf": 0.34
			},
			"weather": [
			  {
				"id": 800,
				"main": "Clear",
				"description": "clear sky",
				"icon": "01n"
			  }
			],
			"clouds": {
			  "all": 0
			},
			"wind": {
			  "speed": 3.82,
			  "deg": 84,
			  "gust": 6.53
			},
			"visibility": 10000,
			"pop": 0,
			"sys": {
			  "pod": "n"
			},
			"dt_txt": "2022-03-15 12:00:00"
		  },
		  {
			"dt": 1647356400,
			"main": {
			  "temp": 285.55,
			  "feels_like": 284.15,
			  "temp_min": 283.34,
			  "temp_max": 285.55,
			  "pressure": 1009,
			  "sea_level": 1009,
			  "grnd_level": 984,
			  "humidity": 50,
			  "temp_kf": 2.21
			},
			"weather": [
			  {
				"id": 800,
				"main": "Clear",
				"description": "clear sky",
				"icon": "01n"
			  }
			],
			"clouds": {
			  "all": 0
			},
			"wind": {
			  "speed": 2.66,
			  "deg": 66,
			  "gust": 4.91
			},
			"visibility": 10000,
			"pop": 0,
			"sys": {
			  "pod": "n"
			},
			"dt_txt": "2022-03-15 15:00:00"
		  },
		  {
			"dt": 1647367200,
			"main": {
			  "temp": 283.98,
			  "feels_like": 282.84,
			  "temp_min": 282.64,
			  "temp_max": 283.98,
			  "pressure": 1011,
			  "sea_level": 1011,
			  "grnd_level": 985,
			  "humidity": 66,
			  "temp_kf": 1.34
			},
			"weather": [
			  {
				"id": 801,
				"main": "Clouds",
				"description": "few clouds",
				"icon": "02n"
			  }
			],
			"clouds": {
			  "all": 12
			},
			"wind": {
			  "speed": 0.81,
			  "deg": 75,
			  "gust": 1.14
			},
			"visibility": 10000,
			"pop": 0,
			"sys": {
			  "pod": "n"
			},
			"dt_txt": "2022-03-15 18:00:00"
		  },
		{
			"dt": 1647766800,
			"main": {
			  "temp": 281.23,
			  "feels_like": 279.99,
			  "temp_min": 281.23,
			  "temp_max": 281.23,
			  "pressure": 1012,
			  "sea_level": 1012,
			  "grnd_level": 983,
			  "humidity": 78,
			  "temp_kf": 0
			},
			"weather": [
			  {
				"id": 803,
				"main": "Clouds",
				"description": "broken clouds",
				"icon": "04n"
			  }
			],
			"clouds": {
			  "all": 74
			},
			"wind": {
			  "speed": 2.13,
			  "deg": 140,
			  "gust": 3.3
			},
			"visibility": 10000,
			"pop": 0,
			"sys": {
			  "pod": "n"
			},
			"dt_txt": "2022-03-20 09:00:00"
		  }
		],
		"city": {
		  "id": 1851632,
		  "name": "Shuzenji",
		  "coord": {
			"lat": 35,
			"lon": 139
		  },
		  "country": "JP",
		  "population": 0,
		  "timezone": 32400,
		  "sunrise": 1647291333,
		  "sunset": 1647334243
		}
	  }`

	model := ForecastData{}
	if err := json.NewDecoder(strings.NewReader(forecastDataJson)).Decode(&model); err != nil {
		t.Error(err)
	}
}

func TestListGeocodingDataModel(t *testing.T) {
	listGeocodingDataJson := `[
		{
		  "name": "London",
		  "local_names": {
			"af": "Londen",
			"ar": "لندن",
			"ascii": "London",
			"az": "London",
			"bg": "Лондон",
			"ca": "Londres",
			"da": "London",
			"de": "London",
			"el": "Λονδίνο",
			"en": "London",
			"eu": "Londres",
			"fa": "لندن",
			"feature_name": "London",
			"fi": "Lontoo",
			"fr": "Londres",
			"gl": "Londres",
			"he": "לונדון",
			"hi": "लंदन",
			"hr": "London",
			"hu": "London",
			"id": "London",
			"it": "Londra",
			"ja": "ロンドン",
			"la": "Londinium",
			"lt": "Londonas",
			"mk": "Лондон",
			"nl": "Londen",
			"no": "London",
			"pl": "Londyn",
			"pt": "Londres",
			"ro": "Londra",
			"ru": "Лондон",
			"sk": "Londýn",
			"sl": "London",
			"sr": "Лондон",
			"th": "ลอนดอน",
			"tr": "Londra",
			"vi": "Luân Đôn",
			"zu": "ILondon"
		  },
		  "lat": 51.5085,
		  "lon": -0.1257,
		  "country": "GB"
		},
		{
		  "name": "London",
		  "local_names": {
			"ar": "لندن",
			"ascii": "London",
			"bg": "Лондон",
			"de": "London",
			"en": "London",
			"fa": "لندن، انتاریو",
			"feature_name": "London",
			"fi": "London",
			"fr": "London",
			"he": "לונדון",
			"ja": "ロンドン",
			"lt": "Londonas",
			"nl": "London",
			"pl": "London",
			"pt": "London",
			"ru": "Лондон",
			"sr": "Лондон"
		  },
		  "lat": 42.9834,
		  "lon": -81.233,
		  "country": "CA"
		},
		{
		  "name": "London",
		  "local_names": {
			"ar": "لندن",
			"ascii": "London",
			"en": "London",
			"fa": "لندن، اوهایو",
			"feature_name": "London",
			"sr": "Ландон"
		  },
		  "lat": 39.8865,
		  "lon": -83.4483,
		  "country": "US",
		  "state": "OH"
		},
		{
		  "name": "London",
		  "local_names": {
			"ar": "لندن",
			"ascii": "London",
			"en": "London",
			"fa": "لندن، کنتاکی",
			"feature_name": "London",
			"sr": "Ландон"
		  },
		  "lat": 37.129,
		  "lon": -84.0833,
		  "country": "US",
		  "state": "KY"
		},
		{
		  "name": "London",
		  "local_names": {
			"ascii": "London",
			"ca": "Londres",
			"en": "London",
			"feature_name": "London"
		  },
		  "lat": 36.4761,
		  "lon": -119.4432,
		  "country": "US",
		  "state": "CA"
		}
	  ]`

	var model ListGeocodingData
	if err := json.NewDecoder(strings.NewReader(listGeocodingDataJson)).Decode(&model); err != nil {
		t.Error(err)
	}
}

func TestGeocodingByZipGeocodingDataModel(t *testing.T) {
	zipGeocodingDataJson := `{
		"zip": "90210",
		"name": "Beverly Hills",
		"lat": 34.0901,
		"lon": -118.4065,
		"country": "US"
	  }`

	model := ZipGeocodingData{}
	if err := json.NewDecoder(strings.NewReader(zipGeocodingDataJson)).Decode(&model); err != nil {
		t.Error(err)
	}
}

func TestCustomInt(t *testing.T) {
	data := []struct {
		name   string
		value  string
		msgErr string
	}{
		{"type int", `{"cod": 10}`, ""},
		{"type string", `{"cod": "10"}`, ""},
		{"unknown value", `{"cod": [10, 11]}`, "customInt: parsing unknown value"},
	}

	type exampleStruct struct {
		Value customInt `json:"cod"`
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			example := &exampleStruct{}
			err := json.NewDecoder(strings.NewReader(d.value)).Decode(&example)

			var msgErr string
			if err != nil {
				msgErr = err.Error()
			}

			if msgErr != d.msgErr {
				t.Errorf("Expected error message `%s`, got `%s`",
					d.msgErr, msgErr)
			}
		})
	}
}

func TestHttpErrorModel(t *testing.T) {
	data := []struct {
		name    string
		httpErr string
		msgErr  string
	}{
		{
			"http code is int",
			`{"cod":401, "message": "Invalid API key. Please see http://openweathermap.org/faq#error401 for more info."}`,
			"401,Invalid API key. Please see http://openweathermap.org/faq#error401 for more info.",
		},
		{
			"http code is string",
			`{"cod":"400","message":"wrong longitude"}`,
			"400,wrong longitude",
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			err := httpErrorFromJson(strings.NewReader(d.httpErr))

			var msgErr string
			if err != nil {
				msgErr = err.Error()
			}

			if msgErr != d.msgErr {
				t.Errorf("Expected error message `%s`, got `%s`",
					d.msgErr, msgErr)
			}
		})
	}
}

func TestUrlIconWeather(t *testing.T) {
	weather := weather{}
	weather.Icon = "01n"

	urlIconWeather := "https://openweathermap.org/img/wn/01n.png"
	url := weather.UrlIconWeather()

	if url != urlIconWeather {
		t.Errorf("Expected url for icon `%s`, got `%s`",
			urlIconWeather, url)
	}
}
