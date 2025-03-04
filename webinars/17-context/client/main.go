package main

import (
	"context"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	httpClient := http.Client{}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
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
