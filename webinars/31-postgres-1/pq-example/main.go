package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://user:password@localhost:5432/otus-go-basic?sslmode=disable")
	if err != nil {
		log.Fatalf("Couldn't open db connection, err: %s", err.Error())
	}

	insertValues(db)
	selectValues(db)
}

func insertValues(db *sql.DB) {
	_, err := db.Exec("insert into users(name) values('John'), ('Patric')")
	if err != nil {
		log.Fatalf("Couldn't insert users, err: %s", err.Error())
	}

	_, err = db.Exec("insert into books(name, user_id) values('Peace and war', 1), ('Idiot', 2)")
	if err != nil {
		log.Fatalf("Couldn't insert books, err: %s", err.Error())
	}
}

func selectValues(db *sql.DB) {
	rows, err := db.Query("select id, name from users")
	if err != nil {
		log.Fatalf("Couldn't insert users, err: %s", err.Error())
	}
	fmt.Println(rows)

	rows, err = db.Query("select id, name, user_id from books")
	if err != nil {
		log.Fatalf("Couldn't insert books, err: %s", err.Error())
	}
}
