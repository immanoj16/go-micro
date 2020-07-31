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

// NewProduct Creates a new product instance
func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
