package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Group represents a group of people waiting for a car.
type Group struct {
	Id     int   `json:"id"`
	People uint8 `json:"people"`
}

// addGroupHandler handles http request for adding new group to waiting list.
func addGroupHandler(w http.ResponseWriter, r *http.Request) {
	var group *Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	waitingGroups = append(waitingGroups, group)
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

	// Check if the group exist
	index := -1
	for i, group := range waitingGroups {
		if group.Id == id {
			index = i
			break
		}
	}

	// Group not found
	if index == -1 {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Remove group
	waitingGroups = append(waitingGroups[:index], waitingGroups[index+1:]...)
	w.WriteHeader(http.StatusNoContent)
}
