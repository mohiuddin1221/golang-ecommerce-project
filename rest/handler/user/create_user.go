package user

import (
	"ecommerce/databse"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser databse.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		utils.SendError(w, "Pleaser provide valis data", http.StatusBadRequest)
		return
	}
	createUser := newUser.Store()
	utils.SendData(w, createUser, http.StatusCreated)

}
