package signup

import (
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"github.com/coldume/pulse/geetest"
	"github.com/coldume/pulse/session"
)

var emailRegexp = regexp.MustCompile(`\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}`)

func emailExistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	email := r.Form.Get("email")
	if validateEmail(email) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if email == "coolwust@gmail.com" {
		writeJSON(w, true)
		return
	}
	writeJSON(w, false)
}

// TODO: validate, sanitize method
type entryForm struct {
	Email   string               `json:"email"`
	Captcha *geetest.UsedCaptcha `json:"captcha"`
}

func entryFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	sess, err := extractSession(r)
	if err != nil {
		writeServerError(w, err)
		return
	}

	failure := &entryFailure{Failure: &Failure{Alerts: make([]int, 0)}}
	if sess == nil || sess.Get("step").(int) != STEP_ENTRY {
		failure.Alerts = append(failure.Alerts, ALERT_SESSION_EXPIRED)
		writeEntryFailure(w, failure)
		return
	}

	form := &entryForm{}
	if json.NewDecoder(r.Body).Decode(form) != nil || form.Captcha == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	failure.Email = validateEmail(form.Email)
	// TODO: Check email existence

	if sess.String("captcha_id") == form.Captcha.CaptchaID {
	} else if form.Captcha.Validate(sess.ID) {
		sess.Set("captcha_id", form.Captcha.CaptchaID)
	} else {
		failure.Captcha = INPUT_STATUS_INCORRECT
	}

	if failure.Email > 0 || failure.Captcha > 0 {
		writeEntryFailure(w, failure)
		return
	}

	// TODO: Send confirmation email

	sess.Replace(map[string]interface{}{"step": STEP_CONFIRMATION, "email": form.Email})
	sess.Touch(SESSION_AGE)
	if err := sessionStore.Update(sess); err != nil {
		writeServerError(w, err)
		return
	}
	writeConfirmationView(w, sess)
}

func validateEmail(email string) int {
	if email == "" {
		return INPUT_STATUS_EMPTY
	} else if !emailRegexp.MatchString(email) {
		return INPUT_STATUS_MALFORMED
	}
	return 0
}

type entryFailure struct {
	*Failure
	Email   int `json:"email"`
	Captcha int `json:"captcha"`
}

func writeEntryFailure(w http.ResponseWriter, failure *entryFailure) {
	writeJSON(w, failure)
}

type entryView struct {
	*View
	Session *session.Session `json:"session,omitempty"`
	Captcha *geetest.Captcha `json:"captcha"`
}

func writeEntryView(w http.ResponseWriter, sess *session.Session) {
	if sess == nil {
		sess = session.NewSession(time.Now().Add(SESSION_AGE))
		sess.Set("step", STEP_ENTRY)
		var err error
		if sess.ID, err = sessionStore.Insert(sess); err != nil {
			writeServerError(w, err)
			return
		}
	}
	writeJSON(w, &entryView{
		View:    &View{Step: STEP_ENTRY},
		Session: sess,
		Captcha: geetest.NewCaptcha(sess.ID),
	})
}
