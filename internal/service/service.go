package service

import (
	"github.com/jkrus/test_echo_http/pkg/model"
	"github.com/jkrus/test_echo_http/pkg/repository"
)

type User interface {
	Create(user model.User) (int, error)
	GetAll() ([]model.User, error)
	GetById(user model.User) (model.User, error)
	Delete(user model.User) error
	Update(input model.User) error
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.Users),
	}
}
