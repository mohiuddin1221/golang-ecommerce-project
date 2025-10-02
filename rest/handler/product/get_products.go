package product

import (
	"net/http"

	"ecommerce/utils"
)

func (h *Handler) Getproducts(w http.ResponseWriter, r *http.Request) {
	product, err := h.productRepo.List()
	if err != nil {
		utils.SendError(w, "Eroor.......", http.StatusInternalServerError)
	}
	utils.SendData(w, product, 200)
}
