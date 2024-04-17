package implementation

import (
	"socialnetwork/domain"
	"sync"
)

type SessionRepoString struct {
	lock         sync.Mutex
	sessionStore map[string]domain.SessionString
}

func RepoInit() *SessionRepoString {
	return &SessionRepoString{
		lock:         sync.Mutex{},
		sessionStore: make(map[string]domain.SessionString),
	}
}

func (s *SessionRepoString) SaveById(id string) (domain.SessionString, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if val, ok := s.sessionStore[id]; ok {
		return val, false
	}
	val := make(map[string]string)
	session := domain.SessionString{
		Id:     id,
		Values: val,
	}
	s.sessionStore[id] = session
	return session, true
}

func (s *SessionRepoString) FindById(id string) (domain.SessionString, bool) {
	if len(s.sessionStore) == 0 {
		return domain.SessionString{}, false
	}
	val, ok := s.sessionStore[id]
	return val, ok
}

func (s *SessionRepoString) DeleteById(id string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.sessionStore, id)
}
