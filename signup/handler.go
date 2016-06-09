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
	RESP_ERR_SESSION_EXPIRED = 0
	RESP_ERR_INVALID_CAPTCHA = 1
	RESP_ERR_EMAIL_EXISTS    = 2
	RESP_ERR_INVALID_EMAIL   = 3
)

var (
	Mux          *http.ServeMux
	sessionStore store.Store
)

var (
	SessionCookieName = "signup_sid"
	SessionCookieAge  = 2 * 24 * 60 * 60
	SessionCookiePath = "/sign-up"
	SessionCookieKey  = []byte("secret")
)

func init() {
	Mux = http.NewServeMux()
	Mux.HandleFunc("/api/sign-up/resolve-view", resolveViewHandler)
	//Mux.HandleFunc("/api/sign-up/submit-email", submitEmailHandler)

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

type ResponseCookie struct {
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
		handleServerError(w, err)
	} else if sess == nil {
		sess = session.NewSession(time.Now().Add(time.Duration(SessionCookieAge)))
		sess.Set("view", VIEW_EMAIL)
		sess.ID, err = sessionStore.Insert(sess)
		if err != nil {
			handleServerError(w, err)
			return
		}
		handleEmailView(w, sess)
	} else {
		handleUnknownView(w, sess)
	}
}

//func submitEmailHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != "POST" {
//		w.WriteHeader(http.StatusMethodNotAllowed)
//		return
//	}
//	sess, err := extract(r)
//	if err == ErrNoSession {
//		handleError(w, RESP_ERR_SESSION_EXPIRED)
//		return
//	} else if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	if view, _ := sess.Get("view"); view != VIEW_EMAIL {
//		handleUnknownView(w, sess)
//		return
//	}
//	data := new(SubmitEmailRequest)
//	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	if !data.Captcha.Validate() {
//		handleError(w, RESP_ERR_INVALID_CAPTCHA)
//		return
//	}
//	// TODO: check email format & existence
//	// TODO: send confirmation email
//	sess.Replace(map[string]interface{}{
//		"view": VIEW_CONFIRMATION,
//		"email": data.Email,
//	})
//	if err := sessionStore.Update(sess); err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	handleConfirmationView(w, sess)
//}

func handleUnknownView(w http.ResponseWriter, sess *session.Session) {
	switch sess.Get("view").(string) {
	case VIEW_EMAIL:        handleEmailView(w, sess)
	case VIEW_CONFIRMATION: handleConfirmationView(w, sess)
	case VIEW_ACCOUNT:      handleAccountView(w, sess)
	}
}

func handleEmailView(w http.ResponseWriter, sess *session.Session) {
	handleJSON(w, &EmailViewResponse{
		View:    VIEW_EMAIL,
		Cookie:  &ResponseCookie{
			Name:   SessionCookieName,
			Value:  session.Sign(sess.ID, SessionCookieKey),
			Path:   SessionCookiePath,
			MaxAge: SessionCookieAge,
		},
		Captcha: geetest.NewCaptcha(sess.ID),
	})
}

func handleConfirmationView(w http.ResponseWriter, sess *session.Session) {
	handleJSON(w, &ConfirmationViewResponse{
		View:  VIEW_CONFIRMATION,
		Email: sess.Get("email").(string),
	})
}

func handleAccountView(w http.ResponseWriter, sess *session.Session) {
	handleJSON(w, &AccountViewResponse{
		View:  VIEW_ACCOUNT,
		Email: sess.Get("email").(string),
	})
}

func handleJSON(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		handleServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func handleError(w http.ResponseWriter, errs ...int) {
	handleJSON(w, &ErrorResponse{Errors: errs})
}

func handleServerError(w http.ResponseWriter, err error) {
	// TODO: use better log
	log.Printf("%#v", err)
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
