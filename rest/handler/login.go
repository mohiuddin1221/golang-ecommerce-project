package handler

import (
	"ecommerce/databse"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	var Loginreq LogInRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&Loginreq)
	if err != nil {
		utils.SendError(w, "Pleaser provide valis data", http.StatusBadRequest)
		return
	}
	usr := databse.FindUser(Loginreq.Email, Loginreq.Password)
	if usr == nil {
		utils.SendError(w, "Inavlid emailor password", http.StatusBadRequest)
		return
	}
	utils.SendData(w, usr, http.StatusCreated)

}
