package captcha

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Response struct {
	Success     bool     `json:"success"`
	ChallengeTS string   `json:"challenge_ts,omitempty"`
	Hostname    string   `json:"hostname,omitempty"`
	ErrorCodes  []string `json:"error-codes,omitempty"`
}

func Verify(token string) bool {
	if token == "" {
		return false
	}

	secret := os.Getenv("RECAPTCHA_SECRET_KEY")
	if secret == "" {
		return false
	}

	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify",
		url.Values{
			"secret":   {secret},
			"response": {token},
		})
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return false
	}

	return result.Success
}
