package signup

import (
	"net/http"


	"github.com/coldume/pulse/geetest"
	"github.com/coldume/pulse/session/store"
	"github.com/coldume/pulse/session/store/memory"

	"fmt"

)

const (

	STATE_NEW
	STATE_HAS_SID
	STATE_HAS_SID_NEED_GEETEST
	STATE_NEED_CONFIRMATION
	STATE_
	STATE_EMAIL_ADDRESS = "email-address"

	// Email is sent and waiting for confirmation
	STATE_CONFIRMATION_EMAIL = "confirmation-email"

	// Email is confirmed
	STATE_ACCOUNT_INFORMATION = "account-information"
)

var Mux *http.ServeMux

var sessionStore store.Store

func init() {
	Mux = http.NewServeMux()
	Mux.HandleFunc("/test/", TestHandler)

	sessionStore = memory.NewMemory()
}

func StateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if cookie, err := r.Cookie("signup_sid"); err != nil {
		goto sid
	} else if sess, err := sessionStore.Get(cookie.Value); err == store.NoSuchSession {
		goto sid
	} else if err != nil {
		goto bad
	} else {
		switch v, _ := sess.Get("state"); v {
		case STATE_EMAIL_ADDRESS:
			goto emailAddress
		case STATE_CONFIRMATION_EMAIL:
			goto confirmationEmail
		case STATE_ACCOUNT_INFORMATION:
			goto accountInformation
		}
	}

sid:
	// sid
	// geetest captcha
emailAddr:
	// geetest captcha
confirmationEmail:
	// email
accountInformation:
	// email



		//captcha := geetest.NewCaptcha(sess.id)
		//data := struct {
		//	State   string           `json:"state"`
		//	SID     string           `json:"sid"`
		//	Captcha *geetest.Captcha `json:"captcha"`
		//} {
		//	State:   STATE_EMAIL_ADDRESS,
		//	SID:     sess.id,
		//	Captcha: captcha,
		//}

	// TODO: send appropriate message regarding to the request

	// XXX:
	// No SID found in cookie
	// SID is invalid or expired
	// Email is not set in session
}

//func CaptchaHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != "POST" {
//		return
//	}
//	// TODO: send error if sid is invalid or is expried
//	// TODO: sned appropriate message regarding to the request
//}
//
//func EmailAddressHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != "POST" {
//		return
//	}
//	// TODO:
//}
//
//func ConfirmationEmailHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != "POST" {
//		return
//	}
//	// TODO:
//}
//
//func AccountInformationHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != "POST" {
//		return
//	}
//	// TODO:
//}
//
//
//
//type Session struct {
//	Email           string
//	CaptchaVerified bool
//	EmailConfirmed  bool
//}
//
//type Storage map[string]*Session
//
//func (s Storage) Contain(sid string) (ok bool) {
//	_, ok = s[sid]
//	return
//}
//
//var storage Storage
func TestHandler(w http.ResponseWriter, r *http.Request) {
	captcha := geetest.NewCaptcha("fff")
	fmt.Fprintf(w, "%#v\n", captcha)
}
