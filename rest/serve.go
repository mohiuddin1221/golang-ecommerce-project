package rest

import (
	"ecommerce/config"
	"ecommerce/global_router"
	"ecommerce/rest/handler"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

func Start(cnf config.Config) {
	manager := middleware.Manager{}
	manager.Use(middleware.Logger)
	mux := http.NewServeMux()

	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(handler.Getproducts),
		),
	)
	mux.Handle("POST /createProducts",
		manager.With(
			middleware.Authentication(http.HandlerFunc(handler.CreateProduct)),
		),
	)
	mux.Handle("GET /productId/{id}",
		manager.With(
			http.HandlerFunc(handler.GetproductsById),
		),
	)
	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(handler.UpdateProduct),
		),
	)
	mux.Handle("DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(handler.DeleteProduct),
		),
	)

	//crete user route
	mux.Handle("POST /users",
		manager.With(
			http.HandlerFunc(handler.CreateUser),
		),
	)
	mux.Handle("POST /login",
		manager.With(
			http.HandlerFunc(handler.LogIn),
		),
	)

	handlerWithCORS := global_router.CorsMiddleware(mux)
	addr := ":" + strconv.Itoa(cnf.HttpPort) //type casting int to string
	fmt.Println("Server running on port: ", addr)
	err := http.ListenAndServe(addr, handlerWithCORS)
	if err != nil {
		fmt.Println("something went wrong", err)
	}

}
