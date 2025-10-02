package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetproductsById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pid, err := strconv.Atoi(productId)
	if err != nil {
		utils.SendData(w, "please give me integer number", 404)
		return
	}
	product, err := h.productRepo.Get(pid)
	if err != nil {
		utils.SendError(w, "No data Found", 404)
		return

	}
	if product == nil {
		utils.SendError(w, "No data Found", 404)
		return
	}
	utils.SendData(w, product, 200)

}
