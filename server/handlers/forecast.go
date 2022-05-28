package handlers

import "net/http"

type forecastHandler struct {
}

func newForecastHandler() *forecastHandler {
	return &forecastHandler{}
}

func (f *forecastHandler) forecast(w http.ResponseWriter, r *http.Request) {

}
