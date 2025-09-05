package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"server/models"
	"server/tools/captcha"
	"time"
)

func LoginHandler(db *sql.DB, client string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// CORS / типы
		w.Header().Set("Access-Control-Allow-Origin", client) // Должен быть точным origin, не "*"
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		// Не кешируем чувствительное
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Pragma", "no-cache")

		// Preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, "Ошибка с методом", http.StatusMethodNotAllowed)
			return
		}

		jwtSecret := []byte(os.Getenv("JWT_SECRET"))

		if len(jwtSecret) == 0 {
			http.Error(w, "Проблема с конфигурацией сервера", http.StatusInternalServerError)
			return
		}

		var req models.LoginRequest // должен иметь поля: Email, Password, Token (reCAPTCHA)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Неправильный ввод", http.StatusBadRequest)
			return
		}

		// reCAPTCHA/antibot
		if !captcha.Verify(req.Token) {
			http.Error(w, "Ошибка в reCAPTCHA", http.StatusForbidden)
			return
		}

		var hashedPassword, username, role, email string
		var userID int
		err := db.QueryRow(
			"SELECT id, password, username, role, email FROM users WHERE email = $1",
			req.Email,
		).Scan(&userID, &hashedPassword, &username, &role, &email)

		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
			return
		}
		if err != nil {
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
			return
		}

		now := time.Now()

		// Access Token — короткий срок
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":  userID,
			"username": username,
			"role":     role,
			"iat":      now.Unix(),
			"nbf":      now.Unix(),
			"exp":      now.Add(60 * time.Minute).Unix(),
		})
		accessTokenString, err := accessToken.SignedString(jwtSecret)
		if err != nil {
			http.Error(w, "Ошибка создания access token", http.StatusInternalServerError)
			return
		}

		// Refresh Token — длинный срок
		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":  userID,
			"username": username,
			"role":     role,
			"iat":      now.Unix(),
			"nbf":      now.Unix(),
			"exp":      now.Add(30 * 24 * time.Hour).Unix(),
			// желательно добавить jti и хранить/ротировать в БД для инвалидации, если потом дойдём до ротации
		})
		refreshTokenString, err := refreshToken.SignedString(jwtSecret)
		if err != nil {
			http.Error(w, "Ошибка создания refresh token", http.StatusInternalServerError)
			return
		}

		// Ставим refresh_token в HttpOnly cookie.
		// Для localhost можно без Secure и с SameSite=Lax (т.к. это "same-site", хоть и разный порт).
		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshTokenString,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Secure:   false, // на HTTPS выставь true и SameSite=None
			MaxAge:   int((30 * 24 * time.Hour).Seconds()),
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    accessTokenString,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Secure:   false, // на HTTPS выставь true и SameSite=None
			MaxAge:   1,
		})
	}
}
