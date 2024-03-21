package util

import (
	"encoding/json"
	"net/http"
)

type HandlerWithError func(w http.ResponseWriter, r *http.Request) error

func WriteJSON(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func Make(f HandlerWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusOK, map[string]string{"error": err.Error()})
		}
	}
}
