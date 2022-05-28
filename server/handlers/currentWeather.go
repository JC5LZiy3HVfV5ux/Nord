package handlers

import (
	"net/http"
)

type currentWeatherHandler struct {
}

func newCurrentWeatherHandler() *currentWeatherHandler {
	return &currentWeatherHandler{}
}

func (c *currentWeatherHandler) weather(w http.ResponseWriter, r *http.Request) {

}
