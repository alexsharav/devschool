package tokens

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

func RefreshHandler(client string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", client)
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		refreshCookie, err := r.Cookie("refresh_token")
		if err != nil {
			http.Error(w, "Refresh token not found", http.StatusUnauthorized)
			return
		}

		var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

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
		role, ok2 := claims["role"].(string)
		username, ok3 := claims["username"].(string)
		if !ok1 || !ok2 || !ok3 {
			http.Error(w, "Invalid token data", http.StatusUnauthorized)
			return
		}

		now := time.Now()

		// Создаём новый access token
		newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id":  userID,
			"username": username,
			"role":     role,
			"iat":      now.Unix(),
			"nbf":      now.Unix(),
			"exp":      now.Add(60 * time.Minute).Unix(),
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
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Secure:   false, // на HTTPS выставь true и SameSite=None
			MaxAge:   int((1 * time.Hour).Seconds()),
		})
	}
}
