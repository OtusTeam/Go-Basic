package main

import (
	"context"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
)

func main() {
	client := http.Client{}
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		sendRequest(client, request.Context())
	})

	log.Println("Server is listening")
	log.Fatal(http.ListenAndServe(":8088", nil))
}

func sendRequest(httpClient http.Client, ctx context.Context) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/test", nil)
	if err != nil {
		log.Printf("Error while creating request: %s\n", err.Error())
		return
	}
	request.Header.Set("request-id", uuid.NewString())

	response, err := httpClient.Do(request)
	if err != nil {
		log.Printf("Error while sending request: %s\n", err.Error())
		return
	}

	bytes, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	log.Printf("Response from server: %s\n", string(bytes))
}
