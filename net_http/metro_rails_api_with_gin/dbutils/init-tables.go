package dbutils

import (
	"database/sql"
	"log"
)

// InitializeDB function initializes the db with required tables
func InitializeDB(dbDriver *sql.DB) {

	stmt, err := dbDriver.Prepare(train)
	defer stmt.Close()

	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Println(err)
	}

	stmt, err = dbDriver.Prepare(station)

	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Println(err)
	}

	stmt, err = dbDriver.Prepare(schedule)

	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Println(err)
	}

	log.Println("ALL tables were sucessfully created")
}
