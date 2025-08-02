package main

import (
	"log"
	"net/http"

	"server/db"
	"server/handlers/auth"
)

func main() {
	serverMux := http.NewServeMux()
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: serverMux,
	}
	database, err := db.Connect()

	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	serverMux.HandleFunc("/register", handlers.RegisterHandler(database, server.Addr))

	log.Printf("Backend running %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
