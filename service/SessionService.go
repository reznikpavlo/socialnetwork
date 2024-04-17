package service

import (
	"socialnetwork/domain"
	"socialnetwork/repo/SessionRepo/implementation"
	"sync"
)

type SessionService struct {
	lock        sync.Mutex
	maxLifeTime int64
	repo        *implementation.SessionRepoString
}

func NewSessionService(repoImplementation *implementation.SessionRepoString, maxLifeTime int64) *SessionService {
	return &SessionService{
		repo:        repoImplementation,
		maxLifeTime: maxLifeTime,
		lock:        sync.Mutex{},
	}
}

func (s *SessionService) CheckById(id string) (*domain.SessionString, bool) {
	result, ok := s.repo.FindById(id)
	return &result, ok
}
