package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/JC5LZiy3HVfV5ux/nord/internal/services"
	"github.com/JC5LZiy3HVfV5ux/nord/internal/utils"
	"github.com/JC5LZiy3HVfV5ux/nord/pkg/openweather"
)

type forecastHandler struct {
	utils.Parser
	utils.Response
	service services.Forecast
}

func newForecastHandler(service services.Forecast) *forecastHandler {
	return &forecastHandler{
		service: service,
	}
}

func (f *forecastHandler) forecast(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	cnt, err := f.ParseCnt(query)
	if err != nil {
		f.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	for _, v := range strings.Split(r.URL.RawQuery, "&") {
		switch strings.Split(v, "=")[0] {
		case "lat", "lon":
			lat, lon, err := f.ParseCoordinates(query)
			if err != nil {
				f.HttpErrorResponse(w, http.StatusBadRequest, err)
				return
			}
			if err := openweather.ValidCoordinates(lat, lon); err != nil {
				f.HttpErrorResponse(w, http.StatusBadRequest, err)
				return
			}

			response, err := f.service.ForecastByCoordinates(r.Context(), lat, lon, cnt)
			if err != nil {
				f.HttpErrorResponse(w, f.GetStatusCode(err), err)
				return
			}

			f.HttpResponse(w, http.StatusOK, response)
			return
		case "q":
			q, err := f.ParseCityName(query)
			if err != nil {
				f.HttpErrorResponse(w, http.StatusBadRequest, err)
				return
			}

			response, err := f.service.ForecastByCityName(r.Context(), q, cnt)
			if err != nil {
				f.HttpErrorResponse(w, f.GetStatusCode(err), err)
				return
			}

			f.HttpResponse(w, http.StatusOK, response)
			return
		case "id":
			id, err := f.ParseCityID(query)
			if err != nil {
				f.HttpErrorResponse(w, http.StatusBadRequest, err)
				return
			}

			response, err := f.service.ForecastByCityId(r.Context(), id, cnt)
			if err != nil {
				f.HttpErrorResponse(w, f.GetStatusCode(err), err)
				return
			}

			f.HttpResponse(w, http.StatusOK, response)
			return
		case "zip":
			zip, err := f.ParseZipCode(query)
			if err != nil {
				f.HttpErrorResponse(w, http.StatusBadRequest, err)
				return
			}

			response, err := f.service.ForecastByZip(r.Context(), zip, cnt)
			if err != nil {
				f.HttpErrorResponse(w, f.GetStatusCode(err), err)
				return
			}
			
			f.HttpResponse(w, http.StatusOK, response)
			return
		}
	}

	f.HttpErrorResponse(w, http.StatusBadRequest, errors.New("nothing to geocode"))
}
