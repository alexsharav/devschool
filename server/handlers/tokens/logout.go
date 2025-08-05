package tokens

import (
	"encoding/json"
	"net/http"
)

func LogoutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Удаляем куки, установив MaxAge < 0
		expiredCookie := func(name string) *http.Cookie {
			return &http.Cookie{
				Name:     name,
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
				Secure:   false,
			}
		}

		http.SetCookie(w, expiredCookie("access_token"))
		http.SetCookie(w, expiredCookie("refresh_token"))

		json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
	}
}
