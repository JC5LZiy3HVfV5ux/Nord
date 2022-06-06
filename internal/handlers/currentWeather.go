package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/JC5LZiy3HVfV5ux/nord/internal/services"
	"github.com/JC5LZiy3HVfV5ux/nord/internal/utils"
	"github.com/JC5LZiy3HVfV5ux/nord/pkg/openweather"
)

type currentWeatherHandler struct {
	utils.Parser
	utils.Response
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
			lat, lon, err := c.ParseCoordinates(query)
			if err != nil {
				c.HttpErrorResponse(w, http.StatusBadRequest, err)
				return
			}
			if err := openweather.ValidCoordinates(lat, lon); err != nil {
				c.HttpErrorResponse(w, http.StatusBadRequest, err)
				return
			}

			response, err := c.service.CurrentByCoordinates(r.Context(), lat, lon)
			if err != nil {
				c.HttpErrorResponse(w, c.GetStatusCode(err), err)
				return
			}

			c.HttpResponse(w, http.StatusOK, response)
			return
		case "q":
			q, err := c.ParseCityName(query)
			if err != nil {
				c.HttpErrorResponse(w, http.StatusBadRequest, err)
				return
			}

			response, err := c.service.CurrentByCityName(r.Context(), q)
			if err != nil {
				c.HttpErrorResponse(w, c.GetStatusCode(err), err)
				return
			}

			c.HttpResponse(w, http.StatusOK, response)
			return
		case "id":
			id, err := c.ParseCityID(query)
			if err != nil {
				c.HttpErrorResponse(w, http.StatusBadRequest, err)
				return
			}

			response, err := c.service.CurrentByCityId(r.Context(), id)
			if err != nil {
				c.HttpErrorResponse(w, c.GetStatusCode(err), err)
				return
			}

			c.HttpResponse(w, http.StatusOK, response)
			return
		case "zip":
			zip, err := c.ParseZipCode(query)
			if err != nil {
				c.HttpErrorResponse(w, http.StatusBadRequest, err)
				return
			}

			response, err := c.service.CurrentByZip(r.Context(), zip)
			if err != nil {
				c.HttpErrorResponse(w, c.GetStatusCode(err), err)
				return
			}
			
			c.HttpResponse(w, http.StatusOK, response)
			return
		}
	}

	c.HttpErrorResponse(w, http.StatusBadRequest, errors.New("nothing to geocode"))
}
