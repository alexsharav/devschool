package handlers

import (
	"database/sql"
	_ "encoding/json"
	"net/http"
	_ "time"

	_ "server/models"

	_ "golang.org/x/crypto/bcrypt"
)

func LoginHanlder(db *sql.DB, client string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", client)
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusNoContent)
			return
		}

	}
}
