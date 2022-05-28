package openweather

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDo(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			mode := req.URL.Query().Get("mode")
			if mode == "1" {
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte(`{
					"cod":"400",
					"message":"wrong longitude"
					}`))
				return
			}

			if mode == "2" {
				rw.WriteHeader(http.StatusUnauthorized)
				rw.Write([]byte(`{
					"cod":401, 
					"message": "Invalid API key. Please see http://openweathermap.org/faq#error401 for more info."
					}`))
				return
			}

			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(`{
				"zip": "90210",
				"name": "Beverly Hills",
				"lat": 34.0901,
				"lon": -118.4065,
				"country": "US"
			  }`))
		}))
	defer server.Close()

	sender := newSender(server.Client())

	data := []struct {
		name   string
		method string
		url    string
		model  interface{}
		result interface{}
		msgErr string
	}{
		{
			"nil model",
			http.MethodGet,
			server.URL,
			nil,
			nil,
			"model is nil",
		},
		{
			"bad request",
			http.MethodGet,
			server.URL + "?mode=1",
			&Zip{},
			&Zip{},
			"400,wrong longitude",
		},
		{
			"unauthorized request",
			http.MethodGet,
			server.URL + "?mode=2",
			&Zip{},
			&Zip{},
			"401,Invalid API key. Please see http://openweathermap.org/faq#error401 for more info.",
		},
		{
			"successful request",
			http.MethodGet,
			server.URL,
			&Zip{},
			&Zip{
				Zip:     "90210",
				Name:    "Beverly Hills",
				Lat:     34.0901,
				Lon:     -118.4065,
				Country: "US",
			},
			"",
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			err := sender.do(context.Background(), d.method, d.url, d.model, nil)

			if diff := cmp.Diff(d.result, d.model); diff != "" {
				t.Error(diff)
			}

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
