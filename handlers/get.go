package handlers

import (
	"net/http"

	"github.com/immanoj16/go-micro/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//   200: productsResponse

// GetProducts returns the products from the data store
func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the data store
	prods := data.GetProducts()

	// serialize the list to JSON
	err := data.ToJSON(prods, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
