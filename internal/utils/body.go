package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendResponse(w http.ResponseWriter, code int, response interface{}) {
	responseEncoder := json.NewEncoder(w)
	if err := responseEncoder.Encode(response); err != nil {
		log.Fatalln("Unable to encode to json", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

type ErrorResponse struct {
	errors interface{}
}

func SendError(w http.ResponseWriter, code int, errors interface{}) {
	SendResponse(w, code, ErrorResponse{
		errors,
	})
}
