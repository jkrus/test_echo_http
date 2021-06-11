package service

import (
	"github.com/jkrus/test_echo_http/pkg/model"
	"github.com/jkrus/test_echo_http/pkg/repository"
)

type User interface {
	Create(userId int, user model.User) (int, error)
	GetAll() ([]model.User, error)
	GetById(userId int) (model.User, error)
	Delete(userId int) error
	Update(userId int, input model.User) error
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}
