package router

import (
	"category-service/api/model"

	basehandler "github.com/correooke/MyMeal/common/handler"

	baserouter "github.com/correooke/MyMeal/common/router"
	"github.com/gorilla/mux"
)

func NewRouter(h *basehandler.CommonHandler[model.Category]) *mux.Router {
	r := baserouter.NewRouter()

	baserouter.AddCRUDRoutes[model.Category](r, h, "categories")
	baserouter.AddIsAlive(r)
	// IsAlive returns a simple response indicating the server is running.
	// swagger:route GET /isalive categoryes isAlive
	//
	// Responses:
	//
	//	200: successfulOperation

	// GetCategoryes returns a list of all categoryes.
	// swagger:route GET /categoryes categoryes getCategoryes
	// Responses:
	//
	//	200: successfulOperation
	//	500: internalServerError

	// CreateCategory creates a new category.
	// swagger:route POST /categoryes categoryes createCategory
	// Responses:
	//
	//	200: successfulOperation
	//	400: badRequestError
	//	500: internalServerError

	// UpdateCategory updates a category by its ID.
	// swagger:route PUT /categoryes/{id} categoryes updateCategory
	//
	// Responses:
	//
	//	200: successfulOperation
	//	400: badRequestError
	//	404: notFoundError
	//	500: internalServerError

	// DeleteCategory deletes a category by its ID.
	// swagger:route DELETE /categoryes/{id} categoryes deleteCategory
	//
	// Responses:
	//
	//	200: successfulOperation
	//	404: notFoundError
	//	500: internalServerError
	return r
}
