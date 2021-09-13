package handler

import (
	"io"
	"net/http"
)

// PingHandler is a HandlerFunc for /ping requests
func PingHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple server check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, our other dependency
	// by performing a simple Ping, and include them in the response.
	io.WriteString(w, `Pong`)
}
