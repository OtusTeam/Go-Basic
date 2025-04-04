package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"otus/user-types/json-deserialization/model"
	"reflect"
)

func main() {
	var u model.User
	ageField, _ := reflect.TypeOf(u).FieldByName("Name")
	jsonFieldName := ageField.Tag.Get("json") // "name"
	fmt.Println(jsonFieldName)

	bytes, err := os.ReadFile("data/test.json")
	if err != nil {
		log.Fatal(err)
	}

	users := new([]model.User)
	err = json.Unmarshal(bytes, users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*users)
}
