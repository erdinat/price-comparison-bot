package handler

import (
	"encoding/json"
	"github.com/erdinat/internProjectGolang/internal/model"
	"github.com/erdinat/internProjectGolang/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PriceDiffLogHandler struct {
	Service *service.ProductPriceDiffLogService
}

func NewPriceDiffLogHandler(service *service.ProductPriceDiffLogService) *PriceDiffLogHandler {
	return &PriceDiffLogHandler{
		Service: service,
	}
}

func (h *PriceDiffLogHandler) CreatePriceDiffLog(w http.ResponseWriter, r *http.Request) {
	var log model.ProductPriceDiffLog
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	createdProductPriceDiffLog, err := h.Service.CreateProductPriceDiffLog(&log)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProductPriceDiffLog)

}

func (h *PriceDiffLogHandler) UpdatePriceDiffLog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var log model.ProductPriceDiffLog
	if err := h.Service.UpdateProductPriceDiffLog(&log); err != nil {
		http.Error(w, "Unable to update price diff log", http.StatusInternalServerError)
	}

	log.ID = id
	if err := h.Service.UpdateProductPriceDiffLog(&log); err != nil {
		http.Error(w, "Unable to update price diff log", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(log)
}

func (h *PriceDiffLogHandler) DeletePriceDiffLog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteProductPriceDiffLog(id); err != nil {
		http.Error(w, "Unable to delete price diff log", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Price difference log deleted successfully"))
}

func (h *PriceDiffLogHandler) GetAllPriceDiffLog(w http.ResponseWriter, r *http.Request) {
	product, err := h.Service.GetAllProductPriceDiffLogs()
	if err != nil {
		http.Error(w, "Unable to get price diff logs", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
