package main

import (
	"log"
	"net/http"

	"server/db"
	"server/handlers"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	http.HandleFunc("/register", handlers.RegisterHandler(database))

	log.Println("Backend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
