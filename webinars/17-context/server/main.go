package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		select {
		case <-request.Context().Done():
			log.Printf("Request cancelled with err: %s\n", request.Context().Err())
		case <-time.NewTimer(5 * time.Second).C:
			_, _ = writer.Write([]byte("Response from server"))
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
