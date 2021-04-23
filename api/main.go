package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jotajay/beer_api/api/handlers"
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

	handlers.Handler(r, service)

	fileServer := http.FileServer(http.Dir("./api/static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer)).Methods("GET", "OPTIONS")
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
