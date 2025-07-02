package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"otus/postgres/goose-example/sqlc"
)

func main() {
	ctx := context.Background()
	dsn := "postgresql://user:password@localhost:5432/otus-go-basic?sslmode=disable"
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("Couldn't connect db, err: %s", err.Error())
	}
	defer conn.Close(ctx)

	queries := sqlc.New(conn)
	users, _ := queries.GetAllUsers(ctx)
	log.Printf("Users: %+v", users)

	devicesCount, _ := queries.GetUserDevicesCount(ctx, sqlc.GetUserDevicesCountParams{
		Attributes: []byte(`{"HDD": "1TB SSD"}`),
		Name:       "15",
	})
	log.Printf("devicesCount: %+v", devicesCount)
}
