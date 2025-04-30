package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"otus/postgres/gorm-example/model"
)

func main() {
	dsn := "postgresql://user:password@localhost:5432/otus-go-basic?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Couldn't open db, err: %s", err.Error())
	}

	insertAuthors(db)
	insertUsers(db)
	selectUsers(db)
}

func insertAuthors(db *gorm.DB) {
	db.Create(&model.Author{
		Name: "Lev Tolstoy",
	})
}

func insertUsers(db *gorm.DB) {
	db.Save(&model.User{
		Name: "John",
	})
	db.Save([]model.Book{
		{
			Name:     "War and Peace",
			AuthorId: 1,
		},
		{
			Name:     "Sevastopol stories",
			AuthorId: 1,
		}})
}

func selectUsers(db *gorm.DB) {
	var author model.Author
	err := db.Model(&model.Author{}).
		Joins("join books on books.author_id = authors.id").
		Joins("join users on books.user_id = users.id").
		Find(&author).Error
	if err != nil {
		log.Printf("Error while fetching users: %s", err.Error())
	}
	fmt.Println(author)
}
