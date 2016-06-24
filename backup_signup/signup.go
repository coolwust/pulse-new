package signup

import (
	"encoding/json"
	"log"
	"net/http"
)

var (
	Mux *http.ServeMux = http.NewServeMux()
)

const (
	ALERT_SESSION_EXPIRED = iota + 1
)

const (
	INPUT_STATUS_MALFORMED = iota + 1
	INPUT_STATUS_EMPTY
	INPUT_STATUS_EXISTS
	INPUT_STATUS_INCORRECT
)

const (
	STEP_ENTRY = iota
	STEP_CONFIRMATION
	STEP_DETAIL
)

func init() {
	Mux.HandleFunc("/api/sign-up/guide", guideHandler)
	Mux.HandleFunc("/api/sign-up/entry-form", entryFormHandler)
	Mux.HandleFunc("/api/sign-up/email-exists", emailExistsHandler)
}

type View struct {
	Step int `json:"step"`
}

type Failure struct {
	Alerts []int `json:"alerts,omitempty"`
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		writeServerError(w, err) // TODO: bad request
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func writeServerError(w http.ResponseWriter, err error) {
	log.Printf("%#v", err) // TODO: use better log
	w.WriteHeader(http.StatusInternalServerError)
}
