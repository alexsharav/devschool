package tokens

import (
	"net/http"
)

func LogoutHandler(client string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", client)
		w.Header().Set("Access-Control-Allow-Credentials", "true")

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

		http.SetCookie(w, expiredCookie("refresh_token"))
		http.SetCookie(w, expiredCookie("access_token"))
	}
}
