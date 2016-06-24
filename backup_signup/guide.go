package signup

import (
	"net/http"
)

func guideHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	sess, err := extractSession(r)
	if err != nil {
		writeServerError(w, err)
		return
	}

	if sess == nil {
		writeEntryView(w, nil)
		return
	}

	switch sess.Get("step").(int) {
	case STEP_ENTRY:
		writeEntryView(w, sess)
	case STEP_CONFIRMATION:
		writeConfirmationView(w, sess)
	case STEP_DETAIL:
		writeDetailView(w, sess)
	}
}
