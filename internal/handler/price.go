package handler

import (
	"encoding/json"
	"github.com/erdinat/internProjectGolang/internal/model"
	"github.com/erdinat/internProjectGolang/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PriceDiffHandler struct {
	Service service.ProductPriceDiffService
}

func NewPriceDiffHandler(s service.ProductPriceDiffService) *PriceDiffHandler {
	return &PriceDiffHandler{
		Service: s,
	}
}

func (h *PriceDiffHandler) CreatePriceDiff(w http.ResponseWriter, r *http.Request) {
	var p model.ProductPriceDiff
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	createdProductPriceDiff, err := h.Service.CreateProductPriceDiff(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProductPriceDiff)
}

func (h *ProductHandler) GetAllPriceDiff(w http.ResponseWriter, r *http.Request) {
	products, err := h.Service.GetAllPriceDiff()
	if err != nil {
		http.Error(w, "Unable fetch price differences", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)

}

func (h *PriceDiffHandler) UpdatePriceDiff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalis ID", http.StatusBadRequest)
		return
	}
	var p model.ProductPriceDiff
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	p.ID = id
	if err := h.Service.UpdatePriceDiff(&p); err != nil {
		http.Error(w, "Unable to update price", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func (h *PriceDiffHandler) DeletePriceDiff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalis ID", http.StatusBadRequest)
		return
	}
	if err := h.Service.DeletePriceDiff(id); err != nil {
		http.Error(w, "Unable to delete price", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}
