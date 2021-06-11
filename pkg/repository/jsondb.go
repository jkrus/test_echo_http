package repository

import "os"

func NewJSONDB(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	return f, err
}
