package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/JC5LZiy3HVfV5ux/openweather-cache-server/server/services"
)

type currentWeatherHandler struct {
	parser
	response
	service services.CurrentWeather
}

func newCurrentWeatherHandler(service services.CurrentWeather) *currentWeatherHandler {
	return &currentWeatherHandler{
		service: service,
	}
}

func (c *currentWeatherHandler) weather(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	for _, v := range strings.Split(r.URL.RawQuery, "&") {
		switch strings.Split(v, "=")[0] {
		case "lat", "lon":
			lat, lon, err := c.parseCoordinates(query)
			if err != nil {
				c.httpErrorResponse(w, http.StatusBadRequest, err)
				return
			}
			response, err := c.service.CurrentByCoordinates(r.Context(), lat, lon)
			if err != nil {
				c.httpErrorResponse(w, c.getStatusCode(err), err)
				return
			}
			c.httpResponse(w, http.StatusOK, response)
			return
		case "q":
			q, err := c.parseCityName(query)
			if err != nil {
				c.httpErrorResponse(w, http.StatusBadRequest, err)
				return
			}
			response, err := c.service.CurrentByCityName(r.Context(), q)
			if err != nil {
				c.httpErrorResponse(w, c.getStatusCode(err), err)
				return
			}
			c.httpResponse(w, http.StatusOK, response)
			return
		case "id":
			id, err := c.parseCityID(query)
			if err != nil {
				c.httpErrorResponse(w, http.StatusBadRequest, err)
				return
			}
			response, err := c.service.CurrentByCityId(r.Context(), id)
			if err != nil {
				c.httpErrorResponse(w, c.getStatusCode(err), err)
				return
			}
			c.httpResponse(w, http.StatusOK, response)
			return
		case "zip":
			zip, err := c.parseZipCode(query)
			if err != nil {
				c.httpErrorResponse(w, http.StatusBadRequest, err)
				return
			}
			response, err := c.service.CurrentByZip(r.Context(), zip)
			if err != nil {
				c.httpErrorResponse(w, c.getStatusCode(err), err)
				return
			}
			c.httpResponse(w, http.StatusOK, response)
			return
		}
	}

	c.httpErrorResponse(w, http.StatusBadRequest, errors.New("nothing to geocode"))
}
