package memory

import (
	"sync"

	"github.com/coldume/pulse/session"
	"github.com/coldume/pulse/session/store"
)

type Memory struct {
	Sessions map[string]*session.Session
	mu       sync.RWMutex
}

func NewMemory() store.Store {
	return &Memory{Sessions: make(map[string]*session.Session)}
}

func (mem *Memory) Get(id string, _ ...interface{}) (sess *session.Session, err error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()
	sess, ok := mem.Sessions[id]
	if !ok {
		err = store.ErrNoSuchSession
	}
	return
}

func (mem *Memory) Insert(sess *session.Session, _ ...interface{}) (id string, err error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()
	id, err = store.UUID()
	if err != nil {
		return
	}
	mem.Sessions[id] = &session.Session{
		ID:      id,
		Expires: sess.Expires,
		Data:    sess.All(),
	}
	return
}

func (mem *Memory) Update(sess *session.Session, _ ...interface{}) error {
	mem.mu.Lock()
	defer mem.mu.Unlock()
	s, ok := mem.Sessions[sess.ID]
	if !ok {
		return store.ErrNoSuchSession
	}
	s.Expires = sess.Expires
	s.Replace(sess.All())
	return nil
}

func (mem *Memory) Delete(id string, x ...interface{}) error {
	mem.mu.Lock()
	defer mem.mu.Unlock()
	delete(mem.Sessions, id)
	return nil
}
