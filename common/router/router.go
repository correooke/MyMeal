package router

import (
	"net/http"

	"github.com/correooke/MyMeal/common/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/isalive", handlers.IsAlive).Methods(http.MethodGet)
	// ... otros middlewares o configuraciones comunes ...
	return r
}
