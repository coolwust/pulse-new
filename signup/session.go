package signup

import (
	"net/http"
	"time"

	"github.com/coldume/pulse/session"
	"github.com/coldume/pulse/session/store"
	"github.com/coldume/pulse/session/store/memory"
)

const (
	SESSION_NAME = "signup_sid"
	SESSION_AGE  = 2 * 24 * time.Hour
)

var sessionStore = memory.NewMemory()

func extractSession(r *http.Request) (sess *session.Session, _ error) {
	cookie, err := r.Cookie(SESSION_NAME)
	if err != nil {
		return
	}

	sid, err := session.Unsign(cookie.Value)
	if err != nil {
		return
	}

	sess, err = sessionStore.Get(sid)
	if err == store.ErrNoSuchSession {
		return
	} else if err != nil {
		return nil, err
	}
	return
}
