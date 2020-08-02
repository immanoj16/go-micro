package handlers

import (
	"net/http"

	"github.com/immanoj16/go-micro/data"
)

// AddProduct is used to add a new product
func (p *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("Prod: %#v", &prod)
	data.AddProduct(&prod)
}
