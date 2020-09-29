package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"

	"go-service/internal/model"
	"go-service/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(h *service.UserService) *UserHandler {
	return &UserHandler{h}
}

func (c *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req model.User
	er1 := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	result, err := c.service.Create(&req)
	if err != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}
	response, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var merchant model.User
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

	err = c.service.Update(&merchant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Updated"))
}

func (c *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	err := c.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted"))
}

func (c *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(result)
	w.Write(response)
}

func (c *UserHandler) LoadUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	result, err := c.service.Load(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(result)
	w.Write(response)
}