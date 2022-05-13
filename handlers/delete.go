package handlers

import (
	"net/http"
	"prod-api/data"
)

//swagger:route DELETE /products/{id} products deleteProducts
//it deletes a products from database
//responses:
//	201 : noContent

//it deletes a product from database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.l.Println("Handling DELETE Request")

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "error: product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "error : product not found", http.StatusInternalServerError)
		return
	}
}
