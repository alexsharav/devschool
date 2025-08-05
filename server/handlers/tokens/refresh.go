package tokens

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func RefreshHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		refreshCookie, err := r.Cookie("refresh_token")
		if err != nil {
			http.Error(w, "Refresh token not found", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(refreshCookie.Value, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userID, ok1 := claims["user_id"].(float64)
		email, ok2 := claims["email"].(string)
		username, ok3 := claims["username"].(string)
		if !ok1 || !ok2 || !ok3 {
			http.Error(w, "Invalid token data", http.StatusUnauthorized)
			return
		}

		// Создаём новый access token
		newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":  int(userID),
			"email":    email,
			"username": username,
			"exp":      time.Now().Add(15 * time.Minute).Unix(),
		})

		accessTokenString, err := newAccessToken.SignedString(jwtSecret)
		if err != nil {
			http.Error(w, "Failed to generate new access token", http.StatusInternalServerError)
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

		json.NewEncoder(w).Encode(map[string]string{"message": "Access token refreshed"})
	}
}
