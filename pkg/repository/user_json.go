package repository

import (
	"github.com/jkrus/test_echo_http/pkg/model"
	"os"
)

type UserJSON struct {
	file *os.File
}

func NewUserJSON(file *os.File) *UserJSON {
	return &UserJSON{file: file}
}

func (r *UserJSON) Create(userID int, user model.User) (int, error) {

	return 0, nil
}

func (r *UserJSON) GetAll() ([]model.User, error) {

	return nil, nil
}

func (r *UserJSON) GetById(userID int) (model.User, error) {

	var v = model.User{}
	return v, nil
}

func (r *UserJSON) Delete(userID int) error {

	return nil
}

func (r *UserJSON) Update(userID int, user model.User) error {

	return nil
}
