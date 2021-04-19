package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	http.Handle("/", r)

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":3000",
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile)}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server. %v", err)
	}
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
