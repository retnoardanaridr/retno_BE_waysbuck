package routes

import (
	"BE-waysbuck-API/handlers"
	"BE-waysbuck-API/pkg/middleware"
	"BE-waysbuck-API/pkg/mysql"
	"BE-waysbuck-API/repositories"

	"github.com/gorilla/mux"
)

func ToppingRoutes(r *mux.Router) {
	toppingRepository := repositories.RespositoryTopping(mysql.DB)
	h := handlers.HandlerTopping(toppingRepository)

	r.HandleFunc("/toppings", middleware.Auth(h.GetToppings)).Methods("GET")
	r.HandleFunc("/topping/{id}", middleware.Auth(h.GetTopping)).Methods("GET")
	r.HandleFunc("/topping", middleware.Auth(middleware.UploadFile(h.CreateTopping))).Methods("POST")
	r.HandleFunc("/topping/{id}", middleware.Auth(middleware.UploadFile(h.UpdateTopping))).Methods("PATCH")
	r.HandleFunc("/topping/{id}", middleware.Auth(h.DeleteTopping)).Methods("DELETE")

}
