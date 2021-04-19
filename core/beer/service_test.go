package beer_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/jotajay/beer_api/core/beer"
)

var db *sql.DB

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

	db.Exec("DELETE FROM beer")

	s := beer.NewService(db)

	err = s.Store(b)
	if err != nil {
		t.Fatalf("failed to store beer: %v", err)
	}

	beer, err := s.Get(1)
	if beer.Name != "Heineken" {
		t.Fatalf("failed to match name: %v", beer)
	}
	t.Logf("beer stored successfully: %v", b)
}
