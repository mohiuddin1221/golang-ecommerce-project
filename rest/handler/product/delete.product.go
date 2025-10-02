package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pid, err := strconv.Atoi(productId)
	if err != nil {
		println("Please give me valid id")
		return
	}
	err = h.productRepo.Delete(pid)
	if err != nil {
		utils.SendError(w, "something went wrong", http.StatusInternalServerError)
	}
	utils.SendData(w, "succesfully deleted", 200)
}
