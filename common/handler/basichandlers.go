package handler

import (
	"encoding/json"
	"net/http"

	"github.com/correooke/MyMeal/common/model"
	"github.com/correooke/MyMeal/common/service"
	"github.com/gorilla/mux"
)

type CommonHandler[T model.Entity] struct {
	Service service.CommonService[T]
}

func NewCommonHandler[T model.Entity](service service.CommonService[T]) CommonHandler[T] {
	return CommonHandler[T]{Service: service}
}

func (h *CommonHandler[T]) GetAll(w http.ResponseWriter, r *http.Request) {
	entities, err := h.Service.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(entities)
}

func (h *CommonHandler[T]) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	entity, err := h.Service.GetByID(r.Context(), params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(entity)
}

func (h *CommonHandler[T]) Create(w http.ResponseWriter, r *http.Request) {
	var entity T
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

func (h *CommonHandler[T]) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var entity T
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

func (h *CommonHandler[T]) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if err := h.Service.Delete(r.Context(), params["id"]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
