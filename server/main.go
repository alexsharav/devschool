package main

import (
	"log"
	"net/http"

	"server/db"
	"server/handlers/auth"
	"server/handlers/tokens"
)

func main() {
	serverMux := http.NewServeMux()
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: serverMux,
	}
	database, err := db.Connect()
	client := "http://localhost:5173"

	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	serverMux.HandleFunc("/register", auth.RegisterHandler(database, client))
	serverMux.HandleFunc("/login", auth.LoginHandler(database, client))
	serverMux.HandleFunc("/refresh", tokens.RefreshHandler())
	serverMux.HandleFunc("/logout", tokens.LogoutHandler())

	log.Printf("Backend running %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
