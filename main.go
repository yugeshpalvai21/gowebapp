package main

import (
	"fmt"
	"gowebapp/db" // import db to establish database connection
	"gowebapp/routes"
	"log"
	"net/http"
)

func main() {
	db, err := db.InitDB() // initiate database connection

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := routes.SetupRoutes()

	serverAddress := "localhost:3000"
	http.Handle("/", router)

	// Listen and serve
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		fmt.Println("Somethinf went wrong while starting server")
		panic(err)
	} else {
		fmt.Println("SERVER IS RUNNING ON PORT 3000")
	}
}
