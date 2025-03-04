package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT)

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Response from server"))
	})

	server := http.Server{
		Addr: "localhost:8080",
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	<-ch
	log.Printf("Shutting down server")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("Error while shutting down server, %s\n", err.Error())
	}
}
