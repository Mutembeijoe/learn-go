package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mutembeijoe/data_io/net_http/metro_rails_api_with_gin/dbutils"
)

// DB - pointer to db driver
var DB *sql.DB

func getStation(c *gin.Context) {
	var station dbutils.StationResource

	id := c.Param("station-id")

	err := DB.QueryRow("select ID, NAME, CAST(OPENING_TIME as CHAR), CAST(CLOSING_TIME as CHAR) from station WHERE id=?", id).Scan(&station.ID,
		&station.Name, &station.OpeningTime, &station.ClosingTime)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		fmt.Println(station)
		c.JSON(http.StatusOK, gin.H{
			"result": station,
		})
	}
}

func main() {

	DB, err := sql.Open("sqlite3", "./railApi.db")

	if err != nil {
		log.Println("Failed to create db driver: ", err)
	}

	r := gin.Default()

	r.GET("/fetch_station/:station-id", getStation)

	log.Fatal(r.Run())

	dbutils.InitializeDB(DB)
}
