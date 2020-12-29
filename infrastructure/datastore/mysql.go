package datastore

import (
	"database/sql"
	"log"
)

func NewDB(driverName, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
