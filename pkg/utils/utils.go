package utils

import (
	"encoding/json"
	"net/http"
)

// Message returns a message
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond returns a json response
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// SuccessResponse returns Status ok json
func SuccessResponse(fields map[string]interface{}, w http.ResponseWriter) {
	fields["status"] = "success"
	message, err := json.Marshal(fields)
	if err != nil {
		// An error occurred processing the json
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occurred internally"))
	}
	// Send header, status code and output to writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

// ErrorResponse return error json response
func ErrorResponse(statusCode int, error string, w http.ResponseWriter) {
	// Create a new map and fill it
	fields := make(map[string]interface{})
	fields["status"] = "error"
	fields["message"] = error
	message, err := json.Marshal(fields)

	if err != nil {
		// An error occurred processing the json
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occurred internally"))
	}
	// Send header, status code and output to writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(message)

}
