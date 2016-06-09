package signup

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"fmt"

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

const (
	SID_COOKIE_NAME = "signup_sid"
	SID_COOKIE_KEY  = []byte("hello world")
	SID_COOKIE_AGE  = time.Hour * 24 * 2
)

const (
	ERRNO_SESSION_EXPIRED = 0
)

var (
	ErrNoSession = errors.New("No session found from request")
)

var (
)

func init() {
	Mux = http.NewServeMux()
	Mux.HandleFunc("/api/sign-up/resolve-view", resolveViewHandler)
	Mux.HandleFunc("/api/sign-up/submit-email", submitEmailHandler)

	sessionStore = memory.NewMemory()
}

type SubmitEmailRequest struct {
	Email   string               `json:"email"`
	Captcha *geetest.UsedCaptcha `json:"captcha"`
}

type SubmitAccountRequest struct {
	Nickname string `json:"name"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Errors []int `json:"errors"`
}

type ResponseCookie struct {
	Name   string
	Value  string
	Path   string
	MaxAge int
}

type EmailViewResponse struct {
	View    string           `json:view"`
	Cookie  *ResponseCookie  `json:cookie"`
	Captcha *geetest.Captcha `json:"captcha"`
}

type ConfirmationViewResponse struct {
	View  string `json:view"`
	Email string `json:"email"`
}

type AccountViewResponse struct {
	View  string `json:view"`
	Email string `json:"email"`
}

func resolveViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if sess, err := sessionFromRequest(r); err == ErrNoSession {
		snedEmailView(w, "")
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		email, _ := sess.Get("email")
		switch view, _ := sess.Get("view"); view.(string) {
		case VIEW_EMAIL: handleEmailView(w, sess.ID)
		case VIEW_CONFIRMATION: handleConfirmationView(w, email.(string))
		case VIEW_ACCOUNT: snedAccountView(w, email.(string))
		}
	}
}

func submitEmailHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	sess, err := sessionFromRequest(r)
	if err == ErrNoSession {
		handleError(w, ERRNO_SESSION_EXPIRED)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if view, _ := sess.Get("view"); view != VIEW_EMAIL {

	}
	d := new(APIData)
	if err := json.NewDecoder(r.Body).Decode(d); err != nil {
		handleError(w, err)
		return
	}
	// TODO: check json error
	// TODO Send email confirmation
	w.Header().Set("Content-Type", "application/json")
	data := &APIData{View: VIEW_CONFIRMATION, SID: sess.ID, Email: d.Email}
	j, _ := json.Marshal(data) // TODO: error
	w.Write(j)
}

func handleEmailView(w http.ResponseWriter, sid string) {
	if sid == "" {
		sess := session.NewSession(time.Now().Add(SID_COOKIE_AGE))
		sess.Set("view", VIEW_EMAIL)
		var err error
		sid, err = sessionStore.Insert(sess)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	data, err := json.Marshal(&EmailViewResponse{
		View: VIEW_EMAIL,
		SID: session.Sign(sid, SID_COOKIE_KEY),
		Captcha: geetest.NewCaptcha(sid)
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func handleConfirmationView(w http.ResponseWriter, email string) {
	data, err := json.Marshal(&ConfirmationViewResponse{View: VIEW_CONFIRMATION, Email: email})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func handleAccountView(w http.ResponseWriter, email string) {
	data, err := json.Marshal(&AccountViewResponse{View: VIEW_ACCOUNT, Email: email})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func handleError(w http.ResponseWriter, errs ...int) {
	data, err := json.Marshal(&ErrorResponse{Errors: errs})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func sessionFromRequest(r *http.Request) (*session.Session, error) {
	cookie, err := r.Cookie(SID_COOKIE_NAME)
	if err != nil {
		return nil, ErrNoSession
	}

	sid, err := session.Unsign(cookie.Value, SID_COOKIE_KEY)
	if err != nil {
		return nil, ErrNoSession
	}

	sess, err := sessionStore.Get(sid)
	if err == store.ErrNoSuchSession {
		return nil, ErrNoSession
	} else if err != nil {
		return nil, err
	}
	return sess, nil
}
