package handlers

import (
	"log"
	"net/http"

	"github.com/immanoj16/go-micro/data"
)

// Product struct
type Product struct {
	l *log.Logger
}

// NewProduct creates a product handler with the given logger
func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

// ServeHTTP is the main entry point for the handler and satisfies the http.Handler
func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// handle an update

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the data store
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall json", http.StatusBadRequest)
		return
	}

	p.l.Printf("Prod: %#v", prod)
}
