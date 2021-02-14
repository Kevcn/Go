package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kevcn/service/handlers"
)

func main() {

	/// Lesson 1

	// // register a path to the server MUX (Multiplexer that select which handler handles a HTTP request)
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

	// })

	// // Just another end point
	// http.HandleFunc("/goodBye", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("Good bye")
	// })

	// Start and listen the server on a port
	// http.ListenAndServe(":8080", nil)

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	// Create our own server multiplexer
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	// Create a server with timeouts to prevent attacks
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// place on a go func so that it doesn't stop listening on the port
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// making a channel that passes the type Signal from os package
	sigChan := make(chan os.Signal)

	// broadcast an os command into the sigChan channel
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// gracefully shut down
	s.Shutdown(tc)
}
