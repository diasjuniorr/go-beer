package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jotajay/beer_api/core/beer"
)

func GetAllBeers(service beer.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Header.Get("Accept") {
		case "application/json":
			getAllBeersJSON(service, w)
		default:
			getAllBeersHTML(service, w)
		}
	},
	)
}

func getAllBeersJSON(service beer.UseCase, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	beers, err := service.GetAll()
	if err != nil {
		fmt.Printf("failed to get all beers: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(beers)
	return
}

func getAllBeersHTML(service beer.UseCase, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
	ts, err := template.ParseFiles(
		"./api/templates/header.html",
		"./api/templates/index.html",
		"./api/templates/footer.html",
	)

	if err != nil {
		http.Error(w, "Error parsing: "+err.Error(), http.StatusInternalServerError)
	}

	beers, err := service.GetAll()
	if err != nil {
		fmt.Printf("failed to get all beers: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Beers []*beer.Beer
	}{
		Title: "Beers",
		Beers: beers,
	}

	err = ts.Lookup("index.html").ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func PostBeer(service beer.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var beer *beer.Beer

		err := json.NewDecoder(r.Body).Decode(&beer)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(err)
			return
		}

		err = service.Store(beer)
		if err != nil {
			fmt.Printf("error storing beer: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		return
	})
}

func GetBeer(service beer.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id := mux.Vars(r)["id"]
		beerId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Printf("failed to parse id: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		beer, err := service.Get(beerId)
		if err != nil {
			fmt.Printf("failed to get beer: %v", err)
			return
		}

		json.NewEncoder(w).Encode(beer)
		return
	})
}

func RemoveBeer(service beer.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var beer *beer.Beer
		err := json.NewDecoder(r.Body).Decode(&beer)
		if err != nil {
			fmt.Printf("failed to parse beer: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		removed, err := service.Remove(beer)
		if err != nil {
			fmt.Printf("failed to remove beer: %v", err)
			return
		}

		if removed < 1 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		return

	})
}

func UpdateBeer(service beer.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var beer *beer.Beer

		err := json.NewDecoder(r.Body).Decode(&beer)
		if err != nil {
			fmt.Printf("failed to parse beer: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		updated, err := service.Update(beer)
		if err != nil {
			fmt.Printf("failed to updatedBeer beer: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return

		}

		if updated < 1 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		return

	})
}
