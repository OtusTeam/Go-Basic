package main

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"otus/postgres/pgx-example/model"
)

func main() {
	dsn := "postgresql://user:password@localhost:5432/otus-go-basic?sslmode=disable"
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Couldn't open db, err: %s", err.Error())
	}
	db.SetMaxOpenConns(10)
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to db: %s", err.Error())
	}

	insertValues(db)
	selectValues(db)
}

func insertValues(db *sqlx.DB) {
	_, err := db.Exec("insert into users(name) values('John'), ('Patric')")
	if err != nil {
		log.Fatalf("Couldn't insert users, err: %s", err.Error())
	}

	_, err = db.Exec("insert into authors(name) values('Lev Tolstoy'), ('Dostoevsky')")
	if err != nil {
		log.Fatalf("Couldn't insert authors, err: %s", err.Error())
	}

	_, err = db.Exec("insert into books(name, user_id, author_id) values('Peace and war', 1, 1), ('Idiot', 2, 2)")
	if err != nil {
		log.Fatalf("Couldn't insert books, err: %s", err.Error())
	}
}

func selectValues(db *sqlx.DB) {
	rows, err := db.Queryx("select id, name from users")
	if err != nil {
		log.Printf("Couldn't select users, err: %s", err.Error())
	}
	for rows.Next() {
		var user model.User
		err = rows.StructScan(&user)
		if err != nil {
			log.Fatalf("Couldn't parse users, err: %s", err.Error())
		}
		fmt.Println(user)
	}

	rows, err = db.Queryx("select id, name, user_id from books")
	if err != nil {
		log.Printf("Couldn't select books, err: %s", err.Error())
	}

	for rows.Next() {
		var book model.Book
		err = rows.StructScan(&book)
		if err != nil {
			log.Fatalf("Couldn't parse books, err: %s", err.Error())
		}
		fmt.Println(book)
	}
}
