package model

type User struct {
	Id      int    `json:"-"` // Игнорировать в encode/json
	Name    string `json:"name"`
	Age     int    // `json:age`
	friends []int64
}
