package rest

import (
	"ecommerce/config"
	"ecommerce/global_router"
	"ecommerce/rest/handler/product"
	"ecommerce/rest/handler/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}
func (server *Server) Start() {
	manager := middleware.Manager{}
	manager.Use(middleware.Logger)
	mux := http.NewServeMux()

	//User routes
	server.productHandler.RegisterUserRoutes(mux, manager)
	server.userHandler.RegisterUserRoutes(mux, manager)

	//crete user route

	handlerWithCORS := global_router.CorsMiddleware(mux)
	addr := ":" + strconv.Itoa(server.cnf.HttpPort) //type casting int to string
	fmt.Println("Server running on port: ", addr)
	err := http.ListenAndServe(addr, handlerWithCORS)
	if err != nil {
		fmt.Println("something went wrong", err)
	}

}
