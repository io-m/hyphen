package db_connection

import (
	"database/sql"
	"log"
	"os"

	"github.com/io-m/hyphen/pkg/constants"
	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*sql.DB, error) {
	connStr := os.Getenv(constants.POSTGRES_CONNECTION)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	// Ping check to make sure connection is alive
	if err = db.Ping(); err != nil {
		log.Fatalf("error pinging database: %v", err)
	}

	return db, nil
}
