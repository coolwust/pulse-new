package signup

import (
	"net/http"

	"github.com/coldume/pulse/session"
)

type detailView struct {
	*View
	Email string `json:"email"`
}

func writeDetailView(w http.ResponseWriter, sess *session.Session) {
	writeJSON(w, &detailView{
		View:  &View{Step: STEP_DETAIL},
		Email: sess.Get("email").(string),
	})
}
