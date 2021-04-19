package beer_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/jotajay/beer_api/core/beer"
)

func TestStore(t *testing.T) {
	b := &beer.Beer{
		ID:    1,
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}

	db, err := sql.Open("sqlite3", "../../data/beer_test.db")
	if err != nil {
		os.Exit(1)
	}

	err = clearDB(db)
	if err != nil {
		t.Fatalf("error clearing db: %v", err)
	}

	service := beer.NewService(db)

	err = service.Store(b)
	if err != nil {
		t.Fatalf("failed to store beer: %v", err)
	}

	beer, err := service.Get(1)
	if beer.Name != "Heineken" {
		t.Fatalf("failed to match name: %v", beer)
	}

	var newBeerName = "Guiness"

	beer.Name = newBeerName
	err = service.Update(beer)
	if err != nil {
		t.Fatalf("failed updating beer: %v", err)
	}

	updatedBeer, err := service.Get(1)
	if err != nil {
		t.Fatalf("failed getting updatedBeer : %v", updatedBeer)
	}

	if updatedBeer.Name != newBeerName {
		t.Fatalf("failed to update beer name. expected: %v, got: %v", newBeerName, updatedBeer.Name)
	}
}

func clearDB(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM beer")
	if err != nil {
		return err
	}
	return nil
}
