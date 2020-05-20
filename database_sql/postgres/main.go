package main

import (
	"github.com/mutembeijoe/data_io/database_sql/postgres/models"
	"log"
)

func main(){
	_, err:= models.InitDB()

	if err!=nil{
		log.Fatal(err)
	}

	log.Println("Sucessfully created db")
}
