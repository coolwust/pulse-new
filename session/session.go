package session

import (
	"sync"
	"time"
)

type Session struct {
	ID       string
	Expires  time.Time
	Data     map[string]interface{}
	mu       sync.RWMutex
}

func NewSession(expires time.Time) *Session {
	return &Session{
		Expires: expires,
		Data:    make(map[string]interface{}),
	}
}

// For best storage compatibility, only int64, float64, and string types are allowed
func (sess *Session) Set(k string, v interface{}) {
	sess.mu.Lock()
	defer sess.mu.Unlock()
	sess.Data[k] = v
}

func (sess *Session) Get(k string) (v interface{}, ok bool) {
	sess.mu.RLock()
	defer sess.mu.RUnlock()
	v, ok = sess.Data[k]
	return
}

func (sess *Session) Delete(k string) {
	sess.mu.Lock()
	defer sess.mu.Unlock()
	delete(sess.Data, k)
}

func (sess *Session) Replace(Data map[string]interface{}) {
	sess.mu.Lock()
	defer sess.mu.Unlock()
	sess.Data = Data
}

func (sess *Session) All() map[string]interface{} {
	sess.mu.RLock()
	defer sess.mu.RUnlock()
	m := make(map[string]interface{})
	for k, v := range sess.Data {
		m[k] = v
	}
	return m
}

func (sess *Session) Touch(age time.Duration) {
	sess.Expires = time.Now().Add(age)
}
