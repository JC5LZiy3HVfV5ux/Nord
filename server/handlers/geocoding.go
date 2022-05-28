package handlers

import "net/http"

type geocodingHandler struct {
}

func newGeocodingHandler() *geocodingHandler {
	return &geocodingHandler{}
}

func (g *geocodingHandler) direct(w http.ResponseWriter, r *http.Request) {

}
