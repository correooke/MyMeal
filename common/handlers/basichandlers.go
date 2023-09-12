package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/correooke/MyMeal/common/model"
	"github.com/correooke/MyMeal/common/service"
	"github.com/gorilla/mux"
)

type CommonHandler struct {
	Service service.CommonService
}

func (h *CommonHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	entities, err := h.Service.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(entities)
}

func (h *CommonHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	entity, err := h.Service.GetByID(r.Context(), params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(entity)
}

func (h *CommonHandler) Create(w http.ResponseWriter, r *http.Request) {
	var entity model.Entity
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.Service.Create(r.Context(), entity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CommonHandler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var entity model.Entity
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.Service.Update(r.Context(), params["id"], entity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
