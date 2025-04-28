package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	_, err := sql.Open("postgres", "postgresql://localhost:5432/otus-go-basic")
	if err != nil {
		log.Fatalf("Couldn't open db connection, err: %s", err.Error())
	}
}
