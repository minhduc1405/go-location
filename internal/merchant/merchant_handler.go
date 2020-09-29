package merchant

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

type MerchantHandler struct {
	service MerchantService
}

func NewMerchantHandler(h MerchantService) *MerchantHandler {
	return &MerchantHandler{h}
}

func (h *MerchantHandler) CreateMerchant(w http.ResponseWriter, r *http.Request) {
	var req Merchant
	er1 := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.service.Create(&req)
	if err != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}
	response, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *MerchantHandler) UpdateMerchant(w http.ResponseWriter, r *http.Request) {
	var merchant Merchant
	err := json.NewDecoder(r.Body).Decode(&merchant)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := chi.URLParam(r, "id")
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	if id != merchant.Id {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}

	err = h.service.Update(&merchant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Updated"))
}

func (h *MerchantHandler) DeleteMerchant(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	err := h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted"))
}

func (h *MerchantHandler) GetAllMerchants(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(result)
	w.Write(response)
}

func (h *MerchantHandler) LoadMerchant(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	result, err := h.service.Load(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(result)
	w.Write(response)
}
