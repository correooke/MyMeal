package router

import (
	"fmt"
	"net/http"

	"github.com/correooke/MyMeal/common/handler"
	"github.com/correooke/MyMeal/common/model"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	return r
}

func AddIsAlive(r *mux.Router) *mux.Router {
	r.HandleFunc("/isalive", handler.IsAlive).Methods(http.MethodGet)
	// ... otros middlewares o configuraciones comunes ...
	return r
}

func AddCRUDRoutes[T model.Entity](r *mux.Router, ch *handler.CommonHandler[T], baseRoute string) *mux.Router {
	r.HandleFunc(fmt.Sprintf("/%s", baseRoute), ch.GetAll).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("/%s/{id}", baseRoute), ch.GetByID).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("/%s", baseRoute), ch.Create).Methods(http.MethodPost)
	r.HandleFunc(fmt.Sprintf("/%s/{id}", baseRoute), ch.Update).Methods(http.MethodPut)
	r.HandleFunc(fmt.Sprintf("/%s/{id}", baseRoute), ch.Delete).Methods(http.MethodDelete)

	return r
}
