package router

import (
	"dish-service/api/model"
	"net/http"

	basehandler "github.com/correooke/MyMeal/common/handler"
	baserouter "github.com/correooke/MyMeal/common/router"
)

func NewRouter(h *basehandler.CommonHandler[model.Dish]) http.Handler {
	r := baserouter.NewRouter()

	baserouter.AddCRUDRoutes[model.Dish](r, h, "dishes")
	baserouter.AddIsAlive(r)

	return r
}
