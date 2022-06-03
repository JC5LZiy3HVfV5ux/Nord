package handlers

import (
	"log"
	"net/http"

	"github.com/JC5LZiy3HVfV5ux/openweather-cache-server/server/services"
	"github.com/gorilla/mux"
)

func RegisterHandlers(router *mux.Router, services *services.Services) {
	c := newCurrentWeatherHandler(services.CurrentWeather)
	f := newForecastHandler(services.Forecast)

	api := router.PathPrefix("/api/v1").Subrouter()

	ping := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/plain;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("pong")); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	api.HandleFunc("/ping", ping).Methods("GET")
	api.HandleFunc("/weather", c.weather).Methods("GET")
	api.HandleFunc("/forecast", f.forecast).Methods("GET")
}
