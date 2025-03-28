package duckdb

import (
	"database/sql"
	"log"

	_ "github.com/marcboeker/go-duckdb"
)

type DuckDB struct {
	Client *sql.DB
}

func NewDuckDB(path string) (*DuckDB, error) {
	db, err := sql.Open("duckdb", "")
	if err != nil {
		log.Fatal(err)
	}
	return &DuckDB{Client: db}, nil
}
