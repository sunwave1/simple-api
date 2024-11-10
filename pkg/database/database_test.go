package database

import (
	"testing"
)

func TestConnectionDatabase(t *testing.T) {
	db := GetDatabase()
	if db == nil {
		t.Fatalf("Error connecting to database")
	}
	defer db.Close()
}
