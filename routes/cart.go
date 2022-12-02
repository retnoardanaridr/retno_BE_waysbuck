package routes

import (
	"BE-waysbuck-API/handlers"
	"BE-waysbuck-API/pkg/middleware"
	"BE-waysbuck-API/pkg/mysql"
	"BE-waysbuck-API/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/carts", middleware.Auth(h.GetCarts)).Methods("GET")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.CreateCart)).Methods("POST")
	r.HandleFunc("/cart/{id}", h.UpdateCart).Methods(("PATCH"))

}
