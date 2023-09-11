package router

import (
	"category-service/internal/app/category/handler"
	"net/http"

	"github.com/correooke/MyMeal/common/router"
	"github.com/gorilla/mux"
)

func NewRouter(ch *handler.CategoryHandler) *mux.Router {
	r := mux.NewRouter()

	router.AddIsAlive(r)
	r.HandleFunc("/isalive", ch.IsAlive).Methods(http.MethodGet)
	r.HandleFunc("/categories", ch.GetCategories).Methods(http.MethodGet)
	r.HandleFunc("/categories/{id}", ch.GetCategory).Methods(http.MethodGet)
	r.HandleFunc("/categories", ch.CreateCategory).Methods(http.MethodPost)
	r.HandleFunc("/categories/{id}", ch.UpdateCategory).Methods(http.MethodPut)
	r.HandleFunc("/categories/{id}", ch.DeleteCategory).Methods(http.MethodDelete)

	return r
}
