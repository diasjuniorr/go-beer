package beer_test

import (
	"database/sql"
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
		t.Fatalf("failed to initialize db: %v", err)
	}

	s := beer.NewService(db)

	err = s.Store(b)
	if err != nil {
		t.Fatalf("failed to store beer: %v", err)
	}
	t.Logf("beer stored successfully: %v", b)
}
