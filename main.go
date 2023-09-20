package main

import (
	"fmt"
	"gowebapp/db" // import db to establish database connection
	"log"
)

func main() {
	db, err := db.InitDB() // initiate database connection

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Print("Databse Connected Successfully....", db)
}
