package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	dsn := "postgresql://user:password@localhost:5432/otus-go-basic?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
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

func insertValues(db *sql.DB) {
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

func selectValues(db *sql.DB) {
	rows, err := db.Query("select id, name from users")
	if err != nil {
		log.Fatalf("Couldn't insert users, err: %s", err.Error())
	}
	for rows.Next() {
		var id int
		var name string
		if err = rows.Scan(&id, &name); err != nil {
			log.Printf("Error while scanning rows: %s", err.Error())
		}
		log.Printf("User with id: %d, name: %s", id, name)
	}

	rows, err = db.Query("select id, name, user_id from books")
	if err != nil {
		log.Fatalf("Couldn't insert books, err: %s", err.Error())
	}

	for rows.Next() {
		var id int
		var name string
		var userId int
		if err = rows.Scan(&id, &name, &userId); err != nil {
			log.Printf("Error while scanning rows: %s", err.Error())
		}
		log.Printf("Book with id: %d, name: %s, user_id: %d", id, name, userId)
	}
}
