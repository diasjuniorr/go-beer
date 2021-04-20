package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jotajay/beer_api/core/beer"
)

func GetAllBeers(service beer.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		beers, err := service.GetAll()
		if err != nil {
			fmt.Printf("failed to get all beers: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		err = json.NewEncoder(w).Encode(beers)
	},
	)
}

func PostBeer(service beer.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var beer *beer.Beer

		err := json.NewDecoder(r.Body).Decode(&beer)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(err)
		}

		err = service.Store(beer)
		if err != nil {
			fmt.Printf("error storing beer: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusCreated)
	})
}

func GetBeer(service beer.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := mux.Vars(r)["id"]
		beerId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Printf("failed to parse id: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		beer, err := service.Get(beerId)
		if err != nil {
			fmt.Printf("failed to get beer: %v", err)
		}

		json.NewEncoder(w).Encode(beer)
	})
}

func RemoveBeer(service beer.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var beer *beer.Beer
		err := json.NewDecoder(r.Body).Decode(&beer)
		if err != nil {
			fmt.Printf("failed to parse beer: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		err = service.Remove(beer)
		if err != nil {
			fmt.Printf("failed to remove beer: %v", err)
		}

		w.WriteHeader(http.StatusOK)

	})
}
