package main

import (
	"encoding/json"
	"net/http"
)

// Car represents available car.
type Car struct {
	Id    int   `json:"id"`
	Seats uint8 `json:"seats"`
}

// loadCarsHandler handles http request for replace availableCars.
func loadCarsHandler(w http.ResponseWriter, r *http.Request) {
	var newCars []*Car
	err := json.NewDecoder(r.Body).Decode(&newCars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newCars = availableCars
}
