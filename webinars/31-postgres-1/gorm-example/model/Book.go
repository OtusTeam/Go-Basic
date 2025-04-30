package model

type Book struct {
	Id       int
	Name     string
	UserId   *int
	User     *User
	AuthorId int
}
