package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// register a path to the server MUX (Multiplexer that select which handler handles a HTTP request)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")

		// capture the data and potential error from the request body
		d, error := ioutil.ReadAll(r.Body)

		if error != nil {
			// Allows us to send respond status code
			http.Error(rw, "Ooops", http.StatusBadRequest)
			return

		}

		log.Printf("Data %s\n", d)

		fmt.Fprintf(rw, "Hello %s", d)
	})

	// Just another end point
	http.HandleFunc("/goodBye", func(http.ResponseWriter, *http.Request) {
		log.Println("Good bye")
	})

	// Start and listen the server on a port
	http.ListenAndServe(":8080", nil)
}
