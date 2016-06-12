package signup

import (
	"encoding/json"
	"log"
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

const (
	ERROR_SESSION_EXPIRED = 0
	ERROR_INVALID_CAPTCHA = 1
	ERROR_EMAIL_EXISTS    = 2
	ERROR_INVALID_EMAIL   = 3
)

var (
	Mux          *http.ServeMux
	sessionStore store.Store
)

var (
	SessionCookieName = "signup_sid"
	SessionCookieAge  = time.Hour * 24 * 2
	SessionCookiePath = "/sign-up"
	SessionCookieKey  = []byte("secret")
)

func init() {
	Mux = http.NewServeMux()
	Mux.HandleFunc("/api/sign-up/resolve-view", resolveViewHandler)
	Mux.HandleFunc("/api/sign-up/submit-email", submitEmailHandler)

	sessionStore = memory.NewMemory()
}

type SubmitEmailRequest struct {
	Email   string           `json:"email"`
	Captcha *geetest.Captcha `json:"captcha"`
}

type SubmitAccountRequest struct {
	Nickname string `json:"name"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Errors []int `json:"errors"`
}

type ViewResponse struct {
	View string      `json:"view"`
	Data interface{} `json:"data"`
}

type EmailViewData struct {
	Cookie  *Cookie          `json:"cookie,omitempty"`
	Captcha *geetest.Captcha `json:"captcha"`
}

type ConfirmationViewData struct {
	Email string `json:"email"`
}

type AccountViewData struct {
	Email string `json:"email"`
}

type Cookie struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Path   string `json:"path"`
	MaxAge int    `json:"maxAge"`
}

func resolveViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if sess, err := extractSession(r); err != nil {
		handleInternalError(w, err)
	} else if sess == nil {
		sess = session.NewSession(time.Now().Add(SessionCookieAge))
		sess.Set("view", VIEW_EMAIL)
		sess.ID, err = sessionStore.Insert(sess)
		if err != nil {
			handleInternalError(w, err)
			return
		}
		handleEmailView(w, sess)
	} else {
		handleInProgressView(w, sess)
	}
}

func submitEmailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	sess, err := extractSession(r)
	if err != nil {
		handleInternalError(w, err)
		return
	} else if sess == nil {
		handleError(w, ERROR_SESSION_EXPIRED)
		return
	}

	if view := sess.Get("view").(string); view != VIEW_EMAIL {
		handleInProgressView(w, sess)
		return
	}

	data := new(SubmitEmailRequest)
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		handleInternalError(w, err)
		return
	}
	// TODO: check nil captcha

	if !data.Captcha.Validate(sess.ID) {
		handleError(w, ERROR_INVALID_CAPTCHA)
		return
	}

	// TODO: check email format & existence
	// TODO: send confirmation email
	sess.Replace(map[string]interface{}{
		"view":  VIEW_CONFIRMATION,
		"email": data.Email,
	})
	if err := sessionStore.Update(sess); err != nil {
		handleInternalError(w, err)
		return
	}
	handleConfirmationView(w, sess)
}

func handleInProgressView(w http.ResponseWriter, sess *session.Session) {
	switch sess.Get("view").(string) {
	case VIEW_EMAIL:
		handleEmailView(w, sess)
	case VIEW_CONFIRMATION:
		handleConfirmationView(w, sess)
	case VIEW_ACCOUNT:
		handleAccountView(w, sess)
	}
}

func handleEmailView(w http.ResponseWriter, sess *session.Session) {
	handleJSON(w, &ViewResponse{
		View: VIEW_EMAIL,
		Data: &EmailViewData{
			Cookie: &Cookie{
				Value:  session.Sign(sess.ID, SessionCookieKey),
				MaxAge: int(SessionCookieAge.Seconds()),
			},
			Captcha: geetest.NewCaptcha(sess.ID),
		},
	})
}

func handleConfirmationView(w http.ResponseWriter, sess *session.Session) {
	handleJSON(w, &ViewResponse{
		View: VIEW_CONFIRMATION,
		Data: &ConfirmationViewData{
			Email: sess.Get("email").(string),
		},
	})
}

func handleAccountView(w http.ResponseWriter, sess *session.Session) {
	handleJSON(w, &ViewResponse{
		View: VIEW_ACCOUNT,
		Data: &AccountViewData{
			Email: sess.Get("email").(string),
		},
	})
}

func handleJSON(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		handleInternalError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func handleError(w http.ResponseWriter, errs ...int) {
	handleJSON(w, &ErrorResponse{Errors: errs})
}

func handleInternalError(w http.ResponseWriter, err error) {
	log.Printf("%#v", err) // TODO: use better log
	w.WriteHeader(http.StatusInternalServerError)
}

func extractSession(r *http.Request) (*session.Session, error) {
	cookie, err := r.Cookie(SessionCookieName)
	if err != nil {
		return nil, nil
	}

	sid, err := session.Unsign(cookie.Value, SessionCookieKey)
	if err != nil {
		return nil, nil
	}

	sess, err := sessionStore.Get(sid)
	if err == store.ErrNoSuchSession {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return sess, nil
}
