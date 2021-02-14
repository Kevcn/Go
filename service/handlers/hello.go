package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// This is an extension on the hello struct? 	// struct is referenced with a struct
// The 'ServeHTTP' is a method that operates on the struct Hello
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")

	// capture the data and potential error from the request body
	d, error := ioutil.ReadAll(r.Body)

	if error != nil {
		// Allows us to send respond status code
		http.Error(rw, "Ooops", http.StatusBadRequest)
		return

	}

	fmt.Fprintf(rw, "Hello %s", d)
}
