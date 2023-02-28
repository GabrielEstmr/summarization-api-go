package main_utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type responseError struct {
	scopeError string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func ERROR(w http.ResponseWriter, statusCode int, scopeError error) {
	JSON(w, statusCode, responseError{scopeError: scopeError.Error()})
}
