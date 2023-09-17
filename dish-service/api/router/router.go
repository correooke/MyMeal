package router

import (
	"dish-service/api/model"

	basehandler "github.com/correooke/MyMeal/common/handler"
	baserouter "github.com/correooke/MyMeal/common/router"
	"github.com/gorilla/mux"
)

func NewRouter(h *basehandler.CommonHandler[model.Dish]) *mux.Router {
	r := baserouter.NewRouter()

	baserouter.AddCRUDRoutes[model.Dish](r, h, "dishes")
	baserouter.AddIsAlive(r)

	return r
}
