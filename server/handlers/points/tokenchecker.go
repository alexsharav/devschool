// server/api/auth.go
package points

import (
	"encoding/json"
	"net/http"
)

func AccessChecker(client string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", client)
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		_, err := r.Cookie("access_token")
		if err != nil {
			http.Error(w, "access token не найден", http.StatusUnauthorized)
			return
		}

		response := map[string]bool{
			"authenticated": true,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
