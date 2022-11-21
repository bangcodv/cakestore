package repo

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConn() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cake_store?timeout=30s")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(1 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
