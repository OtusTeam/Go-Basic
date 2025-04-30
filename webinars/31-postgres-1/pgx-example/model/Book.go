package model

type Book struct {
	Id     int
	Name   string
	UserId int `db:"user_id"`
}
