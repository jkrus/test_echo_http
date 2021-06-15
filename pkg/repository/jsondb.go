package repository

type JsonDb struct {
	Users UsersJSON
}

func NewJSONDB(path string) (*JsonDb, error) {
	db := new(JsonDb)
	db.Users.Users = make(map[int]string)
	db.Users.Path = path
	err := db.Users.ReadDataUsers(path)
	return db, err
}
