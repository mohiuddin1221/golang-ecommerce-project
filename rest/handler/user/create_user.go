package user

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type userrepo struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser userrepo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		utils.SendError(w, "Pleaser provide valis data", http.StatusBadRequest)
		return
	}
	createUser, err := h.userrepo.Create(repo.User{
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		Email:       newUser.Email,
		Password:    newUser.Password,
		IsShopOwner: newUser.IsShopOwner,
	}) //newUser.Create(newUser)
	if err != nil {
		utils.SendError(w, "something went wrong", http.StatusBadRequest)
		return
	}
	utils.SendData(w, createUser, http.StatusCreated)

}
