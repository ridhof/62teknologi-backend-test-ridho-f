package database

import "testing"

func TestDatabaseConnection(t *testing.T) {
	var _, err = Connect()
	if err != nil {
		t.Fatalf("Database could not connect")
	}
}
