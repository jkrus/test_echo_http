package repository

import (
	"github.com/jkrus/test_echo_http/pkg/model"
)

type Users interface {
	Create(user model.User) (int, error)
	GetAll() ([]model.User, error)
	GetById(user model.User) (model.User, error)
	Delete(user model.User) error
	Update(input model.User) error
}

type Repository struct {
	Users
}

func NewRepository(db *JsonDb) *Repository {
	return &Repository{
		Users: NewUsersJSON(db),
	}
}
