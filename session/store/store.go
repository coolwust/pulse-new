package store

import (
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/coldume/pulse/session"
)

var ErrNoSuchSession = errors.New("session: no such session")

type Store interface {

	// Return ErrNoSuchSession when the session ID is not found
	Get(id string, x ...interface{}) (*session.Session, error)

	Insert(sess *session.Session, x ...interface{}) (id string, err error)

	// Return ErrNoSuchSession when the session ID is not found
	Update(sess *session.Session, x ...interface{}) error

	Delete(id string, x ...interface{}) error
}

func UUID() (string, error) {
	s := make([]byte, 16)
	if _, err := rand.Read(s); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", s[:4], s[4:6], s[6:8], s[8:10], s[10:16]), nil
}
