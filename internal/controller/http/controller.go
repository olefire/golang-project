package http

import (
	"encoding/json"
	"mongoGo/internal/models"
	"mongoGo/internal/services"
	"net/http"
)

type Services struct {
	services.UserManagement
}

type Controller struct {
	Services
}

func NewController(s Services) *Controller {
	return &Controller{
		Services: s,
	}
}

func (ctr *Controller) CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	var req models.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctr.CreateUser(ctx, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ctr *Controller) GetUsersEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodGet {
		http.Error(w, "wrong http method", http.StatusNotFound)
		return
	}

	users, err := ctr.GetUsersInfo(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	SuccessArrRespond(users, w)
}
