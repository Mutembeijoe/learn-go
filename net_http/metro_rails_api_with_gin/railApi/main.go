package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mutembeijoe/data_io/net_http/metro_rails_api_with_gin/dbutils"
)

// DB - pointer to db driver
var DB *sql.DB

func main() {

	DB, err := sql.Open("sqlite3", "./railApi.db")

	if err != nil {
		log.Println("Failed to create db driver: ", err)
	}

	dbutils.InitializeDB(DB)
}
