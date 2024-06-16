package utils

import (
	"encoding/json"
	"net/http"
)

type Exception struct {
	Message string `json:"message"`
}

func RespondWithError(w http.ResponseWriter, code int, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	ErrorLogger.Println(message, err)
	json.NewEncoder(w).Encode(Exception{Message: message})
}
