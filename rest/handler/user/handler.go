package user

import (
	"ecommerce/config"
	"ecommerce/repo"
)

type Handler struct {
	cnf      config.Config
	userrepo repo.UserRepo
}

func NewHandler(
	cnf config.Config,
	userrepo repo.UserRepo,
) *Handler {
	return &Handler{
		cnf:      cnf,
		userrepo: userrepo,
	}
}
