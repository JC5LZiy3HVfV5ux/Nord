package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type response struct{}

func (r response) httpErrorResponse(w http.ResponseWriter, code int, err error) {
	body := fmt.Sprintf(`{"cod":%d,"message":"%s"}`, code, err.Error())
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write([]byte(body)); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (r response) httpJsonResponse(w http.ResponseWriter, code int, model interface{}) {
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
