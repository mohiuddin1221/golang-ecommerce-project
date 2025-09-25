package product

import (
	"net/http"

	"ecommerce/databse"
	"ecommerce/utils"
)

func (h *Handler) Getproducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, databse.List(), 200)
}
