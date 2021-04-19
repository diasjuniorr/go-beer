package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// db, err := sql.Open("sqlite3", "../data/beer.db")
	// if err != nil {
	// 	log.Fatalf("db failed: %v", err)
	// }

	// service := beer.NewService(db)

	r := mux.NewRouter()

	r.HandleFunc("/v1/beers", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}).Methods("GET")

	http.ListenAndServe(":3000", r)
}
