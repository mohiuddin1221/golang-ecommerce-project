package handler

import (
	"ecommerce/databse"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pid, err := strconv.Atoi(productId)
	if err != nil {
		println("Please give me valid id")
		return
	}
	databse.Delete(pid)
	utils.SendData(w, "succesfully deleted", 200)
}
