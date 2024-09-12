package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func GetEnv(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return value
}

func GetEnvInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return intValue
}

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing Request Body")

	}

	return json.NewDecoder(r.Body).Decode(payload)
}

// func WriteJSON(w http.ResponseWriter, status int , v any ) error {
// 	w.Header().Add("Content-Type" , "application/json")
// 	w.WriteHeader(status)

// 	return json.NewEncoder(w).Encode(v)
// }

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(status)

	json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

//http: superfluous response.WriteHeader call from github.com/arhantbararia/ecom_api/utils.WriteJSON (utils.go:23)
//solve this?
//https://stackoverflow.com/questions/52764306/http-superfluous-response-writeheader-call-from-golang
