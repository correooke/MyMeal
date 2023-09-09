package router

import (
	"dish-service/internal/app/dish/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(dh *handler.DishHandler) *mux.Router {
	r := mux.NewRouter()

	// Define tus rutas aquí
	r.HandleFunc("/isalive", dh.IsAlive).Methods(http.MethodGet)
	r.HandleFunc("/dishes", dh.GetDishes).Methods(http.MethodGet)
	r.HandleFunc("/dishes/{id}", dh.GetDish).Methods(http.MethodGet)
	r.HandleFunc("/dishes", dh.CreateDish).Methods(http.MethodPost)
	r.HandleFunc("/dishes/{id}", dh.UpdateDish).Methods(http.MethodPut)
	r.HandleFunc("/dishes/{id}", dh.DeleteDish).Methods(http.MethodDelete)

	// ... puedes agregar otras rutas o middlewares según lo necesites ...

	return r
}
