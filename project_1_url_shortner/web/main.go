package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/mutembeijoe/data_io/project_1_url_shortner/database"
	"log"
)

type application struct {
	DB *sql.DB
}

func main() {
	db, err := database.InitDB()

	app := &application{DB: db}

	if err != nil {
		log.Fatalln("Failed to initialize Db", err)
	}

	r := gin.Default()

	r.POST("/api/v1/short", app.generateShortUrl)

	log.Fatal(r.Run(":8080"))
}
