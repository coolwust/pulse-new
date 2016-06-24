package signup

import (
	"net/http"
)

type Captcha struct {
	CaptchaID string `json"id"`
	GeetestID string `json:id"`
	Mode      int    `json"mode"`
	Key       string `json"key"`
	Hash      string `json:"hash"`
}

type User struct {
	SID       string  `json:"sid"`
	Email     string  `json:"email"`
	Captcha   Captcha `json"captcha"`
	Confirmed bool    `json:"confirmed"`
	Nickname  string  `json:"nickname"`
	Password  string  `json:"password"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		userPostHandler(w, r)
	case "PATCH":
		userPatchHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func userPostHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("signup_sid")
	if err != nil {
		// no cookie, need create
	}
}
