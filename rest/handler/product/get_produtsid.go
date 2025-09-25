package product

import (
	"ecommerce/databse"
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
	product := databse.Get(pid)
	if product == nil {
		utils.SendError(w, "No data Found", 404)
		return
	}
	utils.SendData(w, product, 200)

}
