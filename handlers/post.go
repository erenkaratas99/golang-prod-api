package handlers

import (
	"net/http"
	"prod-api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST request")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

type KeyProduct struct{}
