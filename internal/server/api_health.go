package server

import (
	"net/http"
)

// HealthGet is the health handler (always return 200)
func HealthGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
