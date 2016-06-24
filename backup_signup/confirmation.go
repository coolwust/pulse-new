package signup

import (
	"net/http"

	"github.com/coldume/pulse/session"
)

type confirmationView struct {
	*View
	Email string `json:"email"`
}

func writeConfirmationView(w http.ResponseWriter, sess *session.Session) {
	writeJSON(w, &confirmationView{
		View:  &View{Step: STEP_CONFIRMATION},
		Email: sess.Get("email").(string),
	})
}
