package handler

import (
	"dish-service/api/model"
	"dish-service/internal/app/dish/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type DishHandler struct {
	Service service.DishService
}

func NewDishHandler(s service.DishService) *DishHandler {
	return &DishHandler{Service: s}
}

// IsAlive returns a simple response indicating the server is running.
// swagger:route GET /isalive dishes isAlive
//
// Responses:
//
//	200: successfulOperation
func (h *DishHandler) IsAlive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is alive!"))
}

// GetDishes returns a list of all dishes.
// swagger:route GET /dishes dishes getDishes
// Responses:
//
//	200: successfulOperation
//	500: internalServerError
func (h *DishHandler) GetDishes(w http.ResponseWriter, r *http.Request) {
	dishes, err := h.Service.GetDishes(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(dishes)
}

func (h *DishHandler) GetDish(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dish, err := h.Service.GetDish(r.Context(), params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(dish)
}

// CreateDish creates a new dish.
// swagger:route POST /dishes dishes createDish
// Responses:
//
//	200: successfulOperation
//	400: badRequestError
//	500: internalServerError
func (h *DishHandler) CreateDish(w http.ResponseWriter, r *http.Request) {
	var dish model.Dish
	if err := json.NewDecoder(r.Body).Decode(&dish); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateDish(r.Context(), dish); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// UpdateDish updates a dish by its ID.
// swagger:route PUT /dishes/{id} dishes updateDish
//
// Responses:
//
//	200: successfulOperation
//	400: badRequestError
//	404: notFoundError
//	500: internalServerError
func (h *DishHandler) UpdateDish(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var dish model.Dish
	if err := json.NewDecoder(r.Body).Decode(&dish); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.UpdateDish(r.Context(), params["id"], dish); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteDish deletes a dish by its ID.
// swagger:route DELETE /dishes/{id} dishes deleteDish
//
// Responses:
//
//	200: successfulOperation
//	404: notFoundError
//	500: internalServerError
func (h *DishHandler) DeleteDish(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if err := h.Service.DeleteDish(r.Context(), params["id"]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
