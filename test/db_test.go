package test

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cake_store?timeout=30s")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	defer db.Close()
}
