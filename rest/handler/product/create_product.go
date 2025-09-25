package product

import (
	"ecommerce/databse"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	/*
		1. Take body information from r.Body
		2.Create an instance using Product struct with body information
		3.Append the instance into ProductList

	*/
	var NewProduct databse.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&NewProduct)
	if err != nil {
		http.Error(w, "please provide vlaid json", 400)
		return
	}
	createdProduct := databse.Store(NewProduct)
	utils.SendData(w, createdProduct, 201)

}
