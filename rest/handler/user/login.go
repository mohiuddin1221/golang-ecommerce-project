package user

import (
	"ecommerce/config"
	"ecommerce/databse"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
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
	cnf := config.GetConfig()
	accessToken, err := utils.CreateJwt(cnf.JwtToken, utils.JWTPayload{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})
	if err != nil {
		utils.SendError(w, "Unauthorised user", http.StatusInternalServerError)
		return

	}
	utils.SendData(w, accessToken, http.StatusCreated)

}
