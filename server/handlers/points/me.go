package points

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"server/models"
)

func UserInfoHandler(client string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", client)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")

		accessToken, err := r.Cookie("access_token")
		if err != nil {
			http.Error(w, "access token не найден", http.StatusUnauthorized)
			return
		}

		jwtSecret := []byte(os.Getenv("JWT_SECRET"))
		if len(jwtSecret) == 0 {
			http.Error(w, "Проблема с конфигурацией сервера", http.StatusInternalServerError)
			return
		}

		token, err := jwt.Parse(accessToken.Value, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Невалидный токен", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Невалидный токен", http.StatusUnauthorized)
			return
		}

		var user models.UserInfo
		user.Username = claims["username"].(string)
		user.UserID = claims["user_id"].(int)
		user.Role = claims["role"].(string)

		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, "Ошибка возврата", http.StatusInternalServerError)
			return
		}
	}
}
