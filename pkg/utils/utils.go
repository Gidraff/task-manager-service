package utils

import (
	"encoding/json"
	"log"
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

// Log wrapper returns log messages
func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before")
		defer log.Println("After")
		h.ServeHTTP(w, r) // call original
	})
}

// checkAPIKey wrapper make sure the key `QUERY` has been passed
// and return Unauthorized error
func checkAPIKey(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Query().Get("key")) == 0 {
			http.Error(w, "missing key", http.StatusUnauthorized)
			return // don't call original handler
		}
		h.ServeHTTP(w, r)
	})
}

// MustParams check Params are passed and returns a BadRequest status
func MustParams(h http.Handler, params ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		for _, param := range params {
			if len(q.Get(param)) == 0 {
				http.Error(w, "missing "+param, http.StatusBadRequest)
				return // exit early
			}
		}
		h.ServeHTTP(w, r) // all params present, proceed
	})
}

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
