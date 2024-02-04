package router

import (
	"menu-service/api/model"
	"net/http"

	basehandler "github.com/correooke/MyMeal/common/handler"
	baserouter "github.com/correooke/MyMeal/common/router"
)

func NewRouter(h *basehandler.CommonHandler[model.Menu]) http.Handler {
	r := baserouter.NewRouter()

	baserouter.AddCRUDRoutes[model.Menu](r, h, "menues")
	baserouter.AddIsAlive(r)
	// IsAlive returns a simple response indicating the server is running.
	// swagger:route GET /isalive menues isAlive
	//
	// Responses:
	//
	//	200: successfulOperation

	// GetMenues returns a list of all menues.
	// swagger:route GET /menues menues getMenues
	// Responses:
	//
	//	200: successfulOperation
	//	500: internalServerError

	// CreateMenu creates a new menu.
	// swagger:route POST /menues menues createMenu
	// Responses:
	//
	//	200: successfulOperation
	//	400: badRequestError
	//	500: internalServerError

	// UpdateMenu updates a menu by its ID.
	// swagger:route PUT /menues/{id} menues updateMenu
	//
	// Responses:
	//
	//	200: successfulOperation
	//	400: badRequestError
	//	404: notFoundError
	//	500: internalServerError

	// DeleteMenu deletes a menu by its ID.
	// swagger:route DELETE /menues/{id} menues deleteMenu
	//
	// Responses:
	//
	//	200: successfulOperation
	//	404: notFoundError
	//	500: internalServerError
	return r
}
