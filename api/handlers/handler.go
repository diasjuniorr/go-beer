package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jotajay/beer_api/core/beer"
)

func Handler(r *mux.Router, service *beer.Service) {
	r.HandleFunc("/v1/beers", GetAllBeers(service)).Methods("GET")
	r.HandleFunc("/v1/beers", PostBeer(service)).Methods("POST")
	r.HandleFunc("/v1/beers/{id}", GetBeer(service)).Methods("GET")
	r.HandleFunc("/v1/beers", RemoveBeer(service)).Methods("DELETE")
}
