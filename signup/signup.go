package signup

import (
	"encoding/json"
	"net/http"
	"net/mail"

	"github.com/coldume/pulse/geetest"
)

const (
	STEP_EMAIL_ADDRESS = iota
	STEP_CONFIRMATION_EMAIL
	STEP_ACCOUNT_DETAIL
)

const (
	ERROR_SESSION_EXPIRED = iota
	ERROR_INCORRECT_CAPTCHA
	ERROR_EMAIL_EXISTS
	ERROR_MALFORMED_EMAIL
	ERROR_CORRUPTED_DATA
)

var Mux *http.ServeMux = http.NewServeMux()

func init() {
	Mux.HandleFunc("/api/sign-up/guide", guideHandler)
	Mux.HandleFunc("/api/sign-up/submit-email-address", submitEmailAddressHandler)
}

func guideHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if sess, err := extractSession(r); err != nil {
		writeInternalError(w, err)
	} else {
		writeView(w, sess)
	}
}

type EmailAddressSubmit struct {
	Email   string               `json:"email"`
	Captcha *geetest.UsedCaptcha `json:"captcha"`
}

func submitEmailAddressHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	sess, err := extractSession(r)
	if err != nil {
		writeInternalError(w, err)
		return
	} else if sess == nil {
		writeError(w, ERROR_SESSION_EXPIRED)
		return
	}

	if step := sess.Get("step").(int); step != STEP_EMAIL_ADDRESS {
		writeView(w, sess)
		return
	}

	submit := &EmailAddressSubmit{}
	if err := json.NewDecoder(r.Body).Decode(submit); err != nil {
		writeInternalError(w, err)
		return
	}

	addr, err := mail.ParseAddress(submit.Email)
	if err != nil {
		writeError(w, ERROR_MALFORMED_EMAIL)
		return
	}

	if submit.Captcha == nil {
		writeError(w, ERROR_CORRUPTED_DATA)
		return
	}

	if !submit.Captcha.Validate(sess.ID) {
		writeError(w, ERROR_INCORRECT_CAPTCHA)
		return
	}

	// TODO: check existence
	// TODO: send confirmation email
	sess.Replace(map[string]interface{}{
		"step":  STEP_CONFIRMATION_EMAIL,
		"email": addr.String(),
	})
	sess.Touch(SESSION_AGE)
	if err := sessionStore.Update(sess); err != nil {
		writeInternalError(w, err)
		return
	}
	writeConfirmationEmailView(w, sess)
}

type AccountDetailSubmit struct {
	Nickname string `json:"name"`
	Password string `json:"password"`
}
