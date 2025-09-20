package handler

import (
	"ecommerce/databse"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pid, err := strconv.Atoi(productId)
	if err != nil {
		utils.SendData(w, "please give me integer number", 404)
		return
	}
	var NewProduct databse.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&NewProduct)
	if err != nil {
		http.Error(w, "please provide vlaid json", 400)
		return
	}
	NewProduct.ID = pid
	databse.Put(NewProduct)
	utils.SendData(w, "Product syccesfully updated", 200)

}
