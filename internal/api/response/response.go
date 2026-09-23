package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Errors string `json:"errors"`
}

func Error(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(Response{Errors: message}); err != nil {
		http.Error(w, "error with encode Error Response", http.StatusInternalServerError)
	}
}
