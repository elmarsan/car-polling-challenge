package main

import "net/http"

// statusHandler handles http request for sharing app status.
func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
