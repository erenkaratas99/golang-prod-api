package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"prod-api/data"
	"strconv"
)

// Package Classification of Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath : /
// Version : 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// all products in the system
	// in : body
	Body []data.Product
}

//swagger:parameters deleteProduct
type productIDWrapper struct {
	// the id of the product to delete it from the DB
	// in: path
	// required : true
	ID int `json:"id"`
}

//swagger:response noContent
type productsNoContent struct {
}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func getProductID(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id
}
