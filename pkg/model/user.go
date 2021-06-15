package model

type User struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
