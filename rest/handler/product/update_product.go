package product

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type UpProduct struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pid, err := strconv.Atoi(productId)
	if err != nil {
		utils.SendData(w, "please give me integer number", 404)
		return
	}
	var upproduct UpProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&upproduct)
	if err != nil {
		http.Error(w, "please provide vlaid json", 400)
		return

	}
	err = h.productRepo.Update(repo.Product{
		ID:          pid,
		Title:       upproduct.Title,
		Description: upproduct.Description,
		Price:       upproduct.Price,
	})
	if err != nil {
		http.Error(w, "please provide vlaid json", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, "Product syccesfully updated", 200)

}
