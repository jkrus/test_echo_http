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

/*file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
defer file.Close()
if err != nil {
	return s, err
}

stat, _ := file.Stat()
data := make([]byte, stat.Size())
file.Read(data)
for{
	_, err := file.Read(data)
	if err == io.EOF{
		break
	}
}
err = json.Unmarshal(data, &s.users)
s.ind = len(s.users)
s.path = path*/
