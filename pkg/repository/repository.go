package repository

import (
	"github.com/jkrus/test_echo_http/pkg/model"
	"os"
)

type User interface {
	Create(userId int, user model.User) (int, error)
	GetAll() ([]model.User, error)
	GetById(userId int) (model.User, error)
	Delete(userId int) error
	Update(userId int, input model.User) error
}

type Repository struct {
	User
	Users *[]model.User
}

func NewRepository(file *os.File) *Repository {
	return &Repository{
		User:  NewUserJSON(file),
		Users: new([]model.User),
	}
}
