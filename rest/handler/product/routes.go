package product

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterUserRoutes(mux *http.ServeMux, manager middleware.Manager) {

	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(h.Getproducts),
		),
	)
	mux.Handle("POST /createProducts",
		manager.With(
			h.middlewares.Authentication(http.HandlerFunc(h.CreateProduct)),
		),
	)
	mux.Handle("GET /productId/{id}",
		manager.With(
			http.HandlerFunc(h.GetproductsById),
		),
	)
	mux.Handle("PUT /products/{id}",
		manager.With(
			h.middlewares.Authentication(http.HandlerFunc(h.UpdateProduct)),
		),
	)
	mux.Handle("DELETE /products/{id}",
		manager.With(
			h.middlewares.Authentication(http.HandlerFunc(h.DeleteProduct)),
		),
	)

}
