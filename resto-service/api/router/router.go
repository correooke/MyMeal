package router

import (
	"resto-service/api/model"
	"net/http"

	basehandler "github.com/correooke/MyMeal/common/handler"

	baserouter "github.com/correooke/MyMeal/common/router"
)

func NewRouter(h *basehandler.CommonHandler[model.Resto]) http.Handler {
	r := baserouter.NewRouter()

	baserouter.AddCRUDRoutes[model.Resto](r, h, "restos")
	baserouter.AddIsAlive(r)
	// IsAlive returns a simple response indicating the server is running.
	// swagger:route GET /isalive restoes isAlive
	//
	// Responses:
	//
	//	200: successfulOperation

	// GetRestoes returns a list of all restoes.
	// swagger:route GET /restoes restoes getRestoes
	// Responses:
	//
	//	200: successfulOperation
	//	500: internalServerError

	// CreateResto creates a new resto.
	// swagger:route POST /restoes restoes createResto
	// Responses:
	//
	//	200: successfulOperation
	//	400: badRequestError
	//	500: internalServerError

	// UpdateResto updates a resto by its ID.
	// swagger:route PUT /restoes/{id} restoes updateResto
	//
	// Responses:
	//
	//	200: successfulOperation
	//	400: badRequestError
	//	404: notFoundError
	//	500: internalServerError

	// DeleteResto deletes a resto by its ID.
	// swagger:route DELETE /restoes/{id} restoes deleteResto
	//
	// Responses:
	//
	//	200: successfulOperation
	//	404: notFoundError
	//	500: internalServerError
	return r
}
