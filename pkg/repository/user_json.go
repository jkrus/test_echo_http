package repository

import (
	"encoding/json"
	"fmt"
	"github.com/jkrus/test_echo_http/pkg/model"
	"io/ioutil"
	"os"
)

type UsersJSON struct {
	Path  string         `json:"path"`
	Ind   int            `json:"ind"`
	Users map[int]string `json:"users"`
}

func NewUsersJSON(db *JsonDb) *UsersJSON {
	return &UsersJSON{db.Users.Path, db.Users.Ind, db.Users.Users}
}

func (r *UsersJSON) ReadDataUsers(path string) error {
	/*file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		return  err
	}
	stat, _ := file.Stat()
	if stat.Size() == 0 {
		return nil
	}
	data := make([]byte, stat.Size())*/

	data, err := ioutil.ReadFile(path)
	//file.Read(data)
	/*for{
		_, err := file.Read(data)
		if err == io.EOF {
			break
		}
	}*/
	if len(data) == 0 {
		return nil
	}
	err = json.Unmarshal(data, &r)
	return err
}

func (r *UsersJSON) WriteDataUsers() error {
	file, err := os.OpenFile(r.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	err = file.Truncate(0)
	data, err := json.Marshal(r)
	file.Write(data)
	return err
}

func (r *UsersJSON) Create(user model.User) (int, error) {
	_, ok := r.Users[user.Id]
	if ok {
		return user.Id, fmt.Errorf("user with ID %d is already exists", user.Id)
	}
	r.Ind++

	r.Users[r.Ind] = user.Name
	f, _ := os.OpenFile(r.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	data, err := json.Marshal(r)
	err = f.Truncate(0)
	f.Write(data)
	return r.Ind, err
}

func (r *UsersJSON) GetAll() ([]model.User, error) {
	err := r.ReadDataUsers(r.Path)
	if err != nil {
		return nil, err
	}
	var users []model.User
	var user model.User
	for key, value := range r.Users {
		user.Id = key
		user.Name = value
		users = append(users, user)
	}
	return users, err
}

func (r *UsersJSON) GetById(user model.User) (model.User, error) {
	err := r.ReadDataUsers(r.Path)
	if err != nil {
		return user, err
	}
	var ok bool
	user.Name, ok = r.Users[user.Id]
	if !ok {
		return user, fmt.Errorf("User with ID %d does not exist", string(user.Id))
	}
	return user, nil
}

func (r *UsersJSON) Update(user model.User) error {
	r.Users[user.Id] = user.Name
	err := r.WriteDataUsers()
	return err
}

func (r *UsersJSON) Delete(user model.User) error {
	delete(r.Users, user.Id)
	err := r.WriteDataUsers()
	return err
}
