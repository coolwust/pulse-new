package signup

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/coldume/pulse/geetest"
	"github.com/coldume/pulse/session"
)

type View struct {
	Step int `json:"step"`
}

func writeView(w http.ResponseWriter, sess *session.Session) {
	if sess == nil {
		writeEmailAddressView(w, sess)
		return
	}
	switch sess.Get("step").(int) {
	case STEP_EMAIL_ADDRESS:
		writeEmailAddressView(w, sess)
	case STEP_CONFIRMATION_EMAIL:
		writeConfirmationEmailView(w, sess)
	case STEP_ACCOUNT_DETAIL:
		writeAccountDetailView(w, sess)
	}
}

type EmailAddressView struct {
	*View
	Session *session.Session `json:"session,omitempty"`
	Captcha *geetest.Captcha `json:"captcha"`
}

func writeEmailAddressView(w http.ResponseWriter, sess *session.Session) {
	if sess == nil {
		sess = session.NewSession(time.Now().Add(SESSION_AGE))
		sess.Set("step", STEP_EMAIL_ADDRESS)
		var err error
		sess.ID, err = sessionStore.Insert(sess)
		if err != nil {
			writeInternalError(w, err)
		}
	}
	writeJSON(w, &EmailAddressView{
		View:    &View{Step: STEP_EMAIL_ADDRESS},
		Session: sess,
		Captcha: geetest.NewCaptcha(sess.ID),
	})
}

type ConfirmationEmailView struct {
	*View
	Email string `json:"email"`
}

func writeConfirmationEmailView(w http.ResponseWriter, sess *session.Session) {
	writeJSON(w, &ConfirmationEmailView{
		View:  &View{Step: STEP_CONFIRMATION_EMAIL},
		Email: sess.Get("email").(string),
	})
}

type AccountDetailView struct {
	*View
	Email string `json:"email"`
}

func writeAccountDetailView(w http.ResponseWriter, sess *session.Session) {
	writeJSON(w, &AccountDetailView{
		View:  &View{Step: STEP_ACCOUNT_DETAIL},
		Email: sess.Get("email").(string),
	})
}

type Error struct {
	Errors []int `json:"errors"`
}

func writeError(w http.ResponseWriter, errs ...int) {
	writeJSON(w, &Error{Errors: errs})
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		writeInternalError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func writeInternalError(w http.ResponseWriter, err error) {
	log.Printf("%#v", err) // TODO: use better log
	w.WriteHeader(http.StatusInternalServerError)
}
