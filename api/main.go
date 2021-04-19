package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jotajay/beer_api/core/beer"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data/beer.db")
	if err != nil {
		log.Fatalf("db failed: %v", err)
	}

	service := beer.NewService(db)

	r := mux.NewRouter()

	r.HandleFunc("/v1/beers", getAll(service)).Methods("GET")

	http.ListenAndServe(":3000", r)
}

func getAll(service beer.UseCase) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("chegou")
		beers, err := service.GetAll()
		if err != nil {
			fmt.Printf("Error getting beer: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		for b, _ := range beers {

			fmt.Printf("beer: %v", b)
		}
	})
}
