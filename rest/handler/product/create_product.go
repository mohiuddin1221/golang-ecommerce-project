package product

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type RequestProduct struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	/*
		1. Take body information from r.Body
		2.Create an instance using Product struct with body information
		3.
		Append the instance into ProductList

	*/
	var req RequestProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "please provide vlaid json", 400)
		return
	}
	createdProduct, err := h.productRepo.Craete(repo.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
	}) //databse.Store(NewProduct)
	if err != nil {
		utils.SendError(w, "Something Went wrong", http.StatusBadRequest)
	}
	utils.SendData(w, createdProduct, 201)

}
