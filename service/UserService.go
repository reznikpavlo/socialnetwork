package service

import (
	"socialnetwork/domain"
	"socialnetwork/repo/SessionRepo/implementation"
	"socialnetwork/repo/userRepo"
	"time"
)

type UserService struct {
	repo    *userRepo.UserRepo
	session *implementation.SessionRepoString
}

func NewUserService(userRepository *userRepo.UserRepo, sessionRepository *implementation.SessionRepoString) *UserService {
	return &UserService{
		repo:    userRepository,
		session: sessionRepository,
	}
}

func (service *UserService) SaveUser(nUser *domain.Usr) *domain.Usr {

	return service.repo.SaveUser(nUser)
}

func (s *UserService) CheckSession(id string) bool {
	_, ok := s.session.FindById(id)
	return ok
}

func (service *UserService) Login(u *domain.Usr) *domain.Usr {
	_, b := service.session.FindById(u.Id)
	if b {
		u.LastVisit = time.Now()
		u.IsActive = true
	} else {
		service.session.SaveById(u.Id)
		u.IsActive = true
	}
	service.repo.SaveUser(u)
	return u
}
func (us *UserService) DeleteSession(id string) {
	us.session.DeleteById(id)
}
func (s *UserService) DeleteUser(u *domain.Usr) {
	usr := s.repo.FindById(u.Id)
	usr.IsActive = false
	_, ok := s.session.FindById(u.Id)
	if ok {
		s.session.DeleteById(u.Id)
	}
	s.repo.DeleteOne(u)
}
