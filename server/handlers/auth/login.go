package auth

import (
	"database/sql"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"server/models"
	"server/tools/captcha"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func LoginHandler(db *sql.DB, client string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", client)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		var req models.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		if !captcha.Verify(req.Token) {
			http.Error(w, "reCAPTCHA verification failed", http.StatusForbidden)
			return
		}

		var hashedPassword, username string
		var userID int

		// Получаем user_id, пароль и username
		err := db.QueryRow("SELECT id, password, username FROM users WHERE email = $1", req.Email).
			Scan(&userID, &hashedPassword, &username)
		if err == sql.ErrNoRows {
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(w, "DB error", http.StatusInternalServerError)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
		if err != nil {
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
			return
		}

		// Access Token: 15 мин
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":  userID,
			"email":    req.Email,
			"username": username,
			"exp":      time.Now().Add(15 * time.Minute).Unix(),
		})

		accessTokenString, err := accessToken.SignedString(jwtSecret)
		if err != nil {
			http.Error(w, "Ошибка создания access token", http.StatusInternalServerError)
			return
		}

		// Refresh Token: 30 дней
		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":  userID,
			"email":    req.Email,
			"username": username,
			"exp":      time.Now().Add(30 * 24 * time.Hour).Unix(),
		})

		refreshTokenString, err := refreshToken.SignedString(jwtSecret)
		if err != nil {
			http.Error(w, "Ошибка создания refresh token", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    accessTokenString,
			Path:     "/",
			Expires:  time.Now().Add(15 * time.Minute),
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Secure:   false,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshTokenString,
			Path:     "/",
			Expires:  time.Now().Add(30 * 24 * time.Hour),
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Secure:   false,
		})

		json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
	}
}
