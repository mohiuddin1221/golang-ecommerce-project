package handler

import (
	"net/http"

	"ecommerce/databse"
	"ecommerce/utils"
)

func Getproducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, databse.List(), 200)
}
