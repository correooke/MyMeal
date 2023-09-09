package handler

import (
	"category-service/api/model"
	"category-service/internal/app/category/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	Service service.CategoryService
}

func NewCategoryHandler(s service.CategoryService) *CategoryHandler {
	return &CategoryHandler{Service: s}
}

// IsAlive returns a simple response indicating the server is running.
// swagger:route GET /isalive categoryes isAlive
//
// Responses:
//
//	200: successfulOperation
func (h *CategoryHandler) IsAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is alive!"))
}

// GetCategoryes returns a list of all categoryes.
// swagger:route GET /categoryes categoryes getCategoryes
// Responses:
//
//	200: successfulOperation
//	500: internalServerError
func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categoryes, err := h.Service.GetCategories(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categoryes)
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	category, err := h.Service.GetCategory(r.Context(), params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

// CreateCategory creates a new category.
// swagger:route POST /categoryes categoryes createCategory
// Responses:
//
//	200: successfulOperation
//	400: badRequestError
//	500: internalServerError
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateCategory(r.Context(), category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// UpdateCategory updates a category by its ID.
// swagger:route PUT /categoryes/{id} categoryes updateCategory
//
// Responses:
//
//	200: successfulOperation
//	400: badRequestError
//	404: notFoundError
//	500: internalServerError
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.UpdateCategory(r.Context(), params["id"], category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteCategory deletes a category by its ID.
// swagger:route DELETE /categoryes/{id} categoryes deleteCategory
//
// Responses:
//
//	200: successfulOperation
//	404: notFoundError
//	500: internalServerError
func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if err := h.Service.DeleteCategory(r.Context(), params["id"]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
