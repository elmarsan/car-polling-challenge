package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// addGroupHandler handles http request for adding new group to waiting list.
func addGroupHandler(w http.ResponseWriter, r *http.Request) {
	var g *Group
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	carPool.addGroup(g)
	w.WriteHeader(http.StatusAccepted)
}

// dropGroupHandler handles http request for remove existing group of waiting list.
func dropGroupHandler(w http.ResponseWriter, r *http.Request) {
	// Check if request has form content type header
	contentTypeHeader := r.Header.Get("Content-Type")
	if contentTypeHeader != "application/x-www-form-urlencoded" {
		http.Error(w, "wrong content type", http.StatusBadRequest)
		return
	}

	// Parse form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract id as string from form values
	idStr := r.FormValue("ID")
	if len(idStr) == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert id to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := carPool.dropGroup(id)

	if res {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

// statusHandler handles http request for sharing app status.
func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// resetStateHandler handles http request for reset app state.
func resetStateHandler(w http.ResponseWriter, r *http.Request) {
	var cars []*Car
	err := json.NewDecoder(r.Body).Decode(&cars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	carPool = NewCarPool(cars)
}
