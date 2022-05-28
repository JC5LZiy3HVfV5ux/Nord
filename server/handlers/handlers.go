package handlers

import (
	"github.com/gorilla/mux"
)

func RegisterHandlers(router *mux.Router) {
	c := newCurrentWeatherHandler()
	f := newForecastHandler()
	g := newGeocodingHandler()

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/weather", c.weather).Methods("GET")
	api.HandleFunc("/forecast", f.forecast).Methods("GET")
	api.HandleFunc("/direct", g.direct).Methods("GET")
}
