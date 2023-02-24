package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	// Available cars
	availableCars = []*Car{}
	// Groups of people waiting for a car
	waitingGroups = []*Group{}
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/status", statusHandler).Methods("GET")
	r.HandleFunc("/cars", loadCarsHandler).Methods("PUT")
	r.HandleFunc("/journey", addGroupHandler).Methods("POST")
	r.HandleFunc("/dropoff", dropGroupHandler).Methods("POST")

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
