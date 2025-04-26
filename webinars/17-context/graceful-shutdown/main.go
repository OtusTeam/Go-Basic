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
	//Package signal will not block sending to c: the caller must ensure that c has sufficient buffer space to keep up
	//with the expected signal rate.
	//For a channel used for notification of just one signal value, a buffer of size 1 is sufficient.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM) // os.Interrupt

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Response from server"))
	})

	server := http.Server{
		Addr: "localhost:8080",
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	s := <-ch
	log.Printf("Receved %v signal, shutting down server\n", s)
	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("Error while shutting down server, %s\n", err.Error())
	}
}
