package handlers

import (
	"log"
	"net/http"

	"github.com/kevcn/service/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.AddProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		// Manually manage getting the ID of the resource from path
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {

	listofProduct := data.GetProducts()

	err := listofProduct.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
		return
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to un-marshal json", http.StatusBadRequest)
		return
	}

	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(prod)
}
