package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/JC5LZiy3HVfV5ux/nord/pkg/openweather"
)

type Response struct{}

func (r Response) HttpErrorResponse(w http.ResponseWriter, code int, err error) {
	response := fmt.Sprintf(`{"cod":%d,"message":"%s"}`, code, err.Error())
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write([]byte(response)); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (r Response) HttpResponse(w http.ResponseWriter, code int, model interface{}) {
	response, err := json.Marshal(model)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(response); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (r Response) GetStatusCode(err error) int {
	var openweatherError *openweather.HttpError
	if errors.As(err, &openweatherError) {
		return int(openweatherError.Cod)
	}
	return http.StatusInternalServerError
}
