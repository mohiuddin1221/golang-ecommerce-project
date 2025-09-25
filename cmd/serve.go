package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handler/product"
	"ecommerce/rest/handler/user"
	"ecommerce/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
	middlewars := middleware.Newmiddleare(cnf)

	productHandler := product.NewHandler(middlewars)
	userHandler := user.NewHandler()

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()

}
