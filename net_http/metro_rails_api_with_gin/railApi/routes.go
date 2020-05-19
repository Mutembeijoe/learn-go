package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/mutembeijoe/data_io/net_http/metro_rails_api_with_gin/dbutils"
)

func getStation(c *gin.Context) {
	var station models.StationResource

	id := c.Param("station-id")

	err := DB.QueryRow("SELECT ID, NAME, CAST(OPENING_TIME as CHAR), CAST(CLOSING_TIME as CHAR) from station WHERE id=?", id).
		Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": station,
		})
	}
}
