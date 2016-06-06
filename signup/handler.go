package signup

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/coldume/pulse/geetest"
	"github.com/coldume/pulse/session"
	"github.com/coldume/pulse/session/store"
	"github.com/coldume/pulse/session/store/memory"
)

const (
	VIEW_EMAIL        = "email"
	VIEW_CONFIRMATION = "confirmation"
	VIEW_ACCOUNT      = "account"
)

var Mux *http.ServeMux

var sessionStore store.Store

func init() {
	Mux = http.NewServeMux()
	Mux.HandleFunc("/sign-up/api/resolve/", ResolveHandler)

	sessionStore = memory.NewMemory()
}


type ViewData struct {
	View    string           `json:"view,omitempty"`
	SID     string           `json:"sid,omitempty"`
	Email   string           `json:"email,omitempty"`
	Captcha *geetest.Captcha `json:"captcha,omitempty"`
}

type ErrorData struct {
	Errors []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `josn:"errors,omitempty"`
}

func ResolveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if sess, err := sessionFromRequest(r); err != nil {
		handleError(w, err)
	} else if sess == nil {
		handleEmail(w, "")
	} else {
		email, _ := sess.Get("email")
		switch view, _ := sess.Get("view"); view.(string) {
		case VIEW_EMAIL: handleEmail(w, sess.ID)
		case VIEW_CONFIRMATION: handleConfirmation(w, email.(string))
		case VIEW_ACCOUNT: handleAccount(w, email.(string))
		}
	}
}

func handleEmail(w http.ResponseWriter, sid string) {
	if sid == "" {
		sess := session.NewSession(time.Now().Add(time.Hour * 24 * 2))
		var err error
		sid, err = sessionStore.Insert(sess)
		if err != nil {
			handleError(w, err)
			return
		}
	}
	data := &ViewData{View: VIEW_EMAIL, SID: sid, Captcha: geetest.NewCaptcha(sid)}
	j, err := json.Marshal(data)
	if err != nil {
		handleError(w, err)
	}
	w.Write(j)
}

func handleError(w http.ResponseWriter, err error) {
}

func handleConfirmation(w http.ResponseWriter, email string) {
}

func handleAccount(w http.ResponseWriter, email string) {
}

func sessionFromRequest(r *http.Request) (*session.Session, error) {
	// TODO: Signature
	cookie, err := r.Cookie("signup_sid")
	if err != nil {
		return nil, nil
	}

	sess, err := sessionStore.Get(cookie.Value)
	if err == store.ErrNoSuchSession {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return sess, nil
}
