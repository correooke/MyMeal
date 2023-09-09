package router

import (
	"net/http"

	"github.com/correooke/MyMeal/tree/main/common/handlers"

	"github.com/gorilla/mux"
)

func AddIsAlive(r *mux.Router) *mux.Router {
	r.HandleFunc("/isalive", handlers.IsAlive).Methods(http.MethodGet)
	// ... otros middlewares o configuraciones comunes ...
	return r
}
