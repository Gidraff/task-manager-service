package helpers

import (
	"encoding/json"
	"net/http"
)

const (
	// ContentType http header
	ContentType = "Content-Type"

	// AppJSON headerValue
	AppJSON = "application/json"
)

// Response return error json response
func Response(statusCode int, msg string, w http.ResponseWriter) {
	// Create a new map and fill it
	fields := make(map[string]interface{})
	fields["message"] = msg
	message, err := json.Marshal(fields)

	if err != nil {
		// An error occurred processing the json
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occurred internally"))
	}
	w.Header().Set(ContentType, AppJSON)
	w.WriteHeader(statusCode)
	w.Write(message)
}

// AuthResponse return tokens
func AuthResponse(statusCode int, at string, w http.ResponseWriter) {
	fields := map[string]string{
		"access_token": at,
	}
	jsonResponse, err := json.Marshal(fields)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occurred internally"))
	}

	w.Header().Set(ContentType, AppJSON)
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

// PayloadResponse returns a payload
func PayloadResponse(payload interface{}, w http.ResponseWriter) {
	w.Header().Set(ContentType, AppJSON)
	json.NewEncoder(w).Encode(payload)
}
