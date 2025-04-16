package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

const address = "localhost:8080"

func writeResponse(writer http.ResponseWriter, _ *http.Request) {
	response := make([]int, 0)
	for i := 0; i < 100_000; i++ {
		response = append(response, i)
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		_, _ = writer.Write([]byte(err.Error()))
	}
	_, _ = writer.Write(bytes)
}

func main() {
	http.HandleFunc("/", writeResponse)
	fmt.Println("Server is listening on ", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
