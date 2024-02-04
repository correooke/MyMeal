package router

import (
	"restaurantGroup-service/api/model"
	"net/http"

	basehandler "github.com/correooke/MyMeal/common/handler"

	baserouter "github.com/correooke/MyMeal/common/router"
)

func NewRouter(h *basehandler.CommonHandler[model.RestaurantGroup]) http.Handler {
	r := baserouter.NewRouter()

	baserouter.AddCRUDRoutes[model.RestaurantGroup](r, h, "restaurantGroups")
	baserouter.AddIsAlive(r)
	// IsAlive returns a simple response indicating the server is running.
	// swagger:route GET /isalive restaurantGroupes isAlive
	//
	// Responses:
	//
	//	200: successfulOperation

	// GetRestaurantGroupes returns a list of all restaurantGroupes.
	// swagger:route GET /restaurantGroupes restaurantGroupes getRestaurantGroupes
	// Responses:
	//
	//	200: successfulOperation
	//	500: internalServerError

	// CreateRestaurantGroup creates a new restaurantGroup.
	// swagger:route POST /restaurantGroupes restaurantGroupes createRestaurantGroup
	// Responses:
	//
	//	200: successfulOperation
	//	400: badRequestError
	//	500: internalServerError

	// UpdateRestaurantGroup updates a restaurantGroup by its ID.
	// swagger:route PUT /restaurantGroupes/{id} restaurantGroupes updateRestaurantGroup
	//
	// Responses:
	//
	//	200: successfulOperation
	//	400: badRequestError
	//	404: notFoundError
	//	500: internalServerError

	// DeleteRestaurantGroup deletes a restaurantGroup by its ID.
	// swagger:route DELETE /restaurantGroupes/{id} restaurantGroupes deleteRestaurantGroup
	//
	// Responses:
	//
	//	200: successfulOperation
	//	404: notFoundError
	//	500: internalServerError
	return r
}
