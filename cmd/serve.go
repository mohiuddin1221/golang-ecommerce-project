package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handler/product"
	"ecommerce/rest/handler/user"
	"ecommerce/rest/middleware"
	"fmt"
	"log"
	"os"
)

func Serve() {
	cnf := config.GetConfig()
	dbcon, err := db.NewConnection()
	dbUrl := ""
	log.Println("Running migrations...")
	db.RunMigration(dbUrl)

	middlewars := middleware.Newmiddleare(cnf)
	productrepo := repo.NewproductRepo(dbcon)
	userrepo := repo.NewUserRepo(dbcon)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productHandler := product.NewHandler(middlewars, productrepo)

	userHandler := user.NewHandler(*cnf, userrepo)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()

}
