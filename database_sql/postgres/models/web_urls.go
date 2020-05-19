package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func InitDB() (*sql.DB, error) {
	connStr := "postgres://orion:goodlife@localhost/goDB?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	stmt, err:= db.Prepare("CREATE TABLE WEB_URL(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")

	if err!=nil{
		return nil, err
	}
	result, err:= stmt.Exec();

	if err!=nil{
		return nil, err
	}

	log.Println(result)

	return db,nil
}
