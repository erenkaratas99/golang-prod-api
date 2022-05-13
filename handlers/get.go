package handlers

import (
	"net/http"
	"prod-api/data"
)

// swagger:route GET /products products listProducts
//it returns a list of products
//responses:
//	200 :productsResponse
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request")
	lp := data.GetProducts() //list of products
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Error : Cannot encode to json", http.StatusInternalServerError)
	}
}
