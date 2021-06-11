package service

import (
	"github.com/jkrus/test_echo_http/pkg/model"
	"github.com/jkrus/test_echo_http/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func (s UserService) Create(userId int, user model.User) (int, error) {

	return 0, nil
}

func (s UserService) GetAll() ([]model.User, error) {

	return nil, nil
}

func (s UserService) GetById(userId int) (model.User, error) {

	var v = model.User{}
	return v, nil
}

func (s UserService) Delete(userId int) error {

	return nil
}

func (s UserService) Update(userId int, input model.User) error {

	return nil
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}
