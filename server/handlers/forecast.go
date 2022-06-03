package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/JC5LZiy3HVfV5ux/openweather-cache-server/server/services"
)

type forecastHandler struct {
	parser
	response
	service services.Forecast
}

func newForecastHandler(service services.Forecast) *forecastHandler {
	return &forecastHandler{
		service: service,
	}
}

func (f *forecastHandler) forecast(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	cnt, err := f.parseCnt(query)
	if err != nil {
		f.httpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	for _, v := range strings.Split(r.URL.RawQuery, "&") {
		switch strings.Split(v, "=")[0] {
		case "lat", "lon":
			lat, lon, err := f.parseCoordinates(query)
			if err != nil {
				f.httpErrorResponse(w, http.StatusBadRequest, err)
				return
			}
			response, err := f.service.ForecastByCoordinates(r.Context(), lat, lon, cnt)
			if err != nil {
				f.httpErrorResponse(w, f.getStatusCode(err), err)
				return
			}
			f.httpResponse(w, http.StatusOK, response)
			return
		case "q":
			q, err := f.parseCityName(query)
			if err != nil {
				f.httpErrorResponse(w, http.StatusBadRequest, err)
				return
			}
			response, err := f.service.ForecastByCityName(r.Context(), q, cnt)
			if err != nil {
				f.httpErrorResponse(w, f.getStatusCode(err), err)
				return
			}
			f.httpResponse(w, http.StatusOK, response)
			return
		case "id":
			id, err := f.parseCityID(query)
			if err != nil {
				f.httpErrorResponse(w, http.StatusBadRequest, err)
				return
			}
			response, err := f.service.ForecastByCityId(r.Context(), id, cnt)
			if err != nil {
				f.httpErrorResponse(w, f.getStatusCode(err), err)
				return
			}
			f.httpResponse(w, http.StatusOK, response)
			return
		case "zip":
			zip, err := f.parseZipCode(query)
			if err != nil {
				f.httpErrorResponse(w, http.StatusBadRequest, err)
				return
			}
			response, err := f.service.ForecastByZip(r.Context(), zip, cnt)
			if err != nil {
				f.httpErrorResponse(w, f.getStatusCode(err), err)
				return
			}
			f.httpResponse(w, http.StatusOK, response)
			return
		}
	}

	f.httpErrorResponse(w, http.StatusBadRequest, errors.New("nothing to geocode"))
}
