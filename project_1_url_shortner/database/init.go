package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://orion:goodlife@localhost/shorturls?sslmode=disable")
	if err != nil {
		return nil, err
	}
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS urls(ID SERIAL, OriginalUrl TEXT NOT NULL )")

	if err != nil {
		return nil, err
	}

	if _, err := stmt.Exec(); err != nil {
		return nil, err
	}
	fmt.Println("Successfully initialised db")
	return db, nil
}
