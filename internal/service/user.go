package service

import (
	"github.com/jkrus/test_echo_http/pkg/model"
	"github.com/jkrus/test_echo_http/pkg/repository"
)

type UserService struct {
	repo repository.Users
}

func (s *UserService) Create(user model.User) (int, error) {
	return s.repo.Create(user)
}

func (s *UserService) GetAll() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetById(user model.User) (model.User, error) {
	return s.repo.GetById(user)
}

func (s *UserService) Delete(user model.User) error {
	return s.repo.Delete(user)
}

func (s *UserService) Update(input model.User) error {
	s.repo.Update(input)
	return nil
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{repo: repo}
}
