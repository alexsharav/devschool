package auth

import (
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"server/models"
	"server/tools/captcha"
)

func RegisterHandler(db *sql.DB, client string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", client)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		var req models.RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Неправильный ввод", http.StatusBadRequest)
			return
		}

		if !captcha.Verify(req.Token) {
			http.Error(w, "Проблема reCAPTCHA", http.StatusForbidden)
			return
		}

		var exists bool
		err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)", req.Email).Scan(&exists)
		if err != nil {
			http.Error(w, "Проблема регистрации", http.StatusInternalServerError)
			return
		}

		if exists {
			http.Error(w, "Проблема регистрации", http.StatusBadRequest)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Проблема регистрации", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
			req.Username, req.Email, string(hash))
		if err != nil {
			http.Error(w, "Проблема регистрации", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
