package http

import (
	"encoding/json"
	"mongoGo/internal/models"
	"mongoGo/internal/services"
	"net/http"
)

type UserService struct {
	services.UserManagement
}

type PasteService struct {
	services.PasteManagement
}

type Controller struct {
	UserService
	PasteService
}

func NewController(us UserService, ps PasteService) *Controller {
	return &Controller{
		UserService:  us,
		PasteService: ps,
	}
}

func (ctr *Controller) CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "wrong http method", http.StatusMethodNotAllowed)
		return
	}

	var req models.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := ctr.CreateUser(ctx, &req)
	if err != nil {
		http.Error(w, "can`t create user", http.StatusInternalServerError)
		return
	}

	payload := struct {
		Id string `json:"id"`
	}{Id: id}
	err = json.NewEncoder(w).Encode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusTeapot)
	}

	w.WriteHeader(http.StatusCreated)
}

func (ctr *Controller) GetUsersEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodGet {
		http.Error(w, "wrong http method", http.StatusMethodNotAllowed)
		return
	}

	users, err := ctr.GetUsers(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SuccessArrRespond(users, w)
}

func (ctr *Controller) CreatePasteEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		http.Error(w, "wrong http method", http.StatusMethodNotAllowed)
		return
	}

	var req models.Paste
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := ctr.CreatePaste(ctx, &req)
	if err != nil {
		http.Error(w, "can`t create paste", http.StatusInternalServerError)
		return
	}

	payload := struct {
		Id string `json:"id"`
	}{Id: id}
	err = json.NewEncoder(w).Encode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusTeapot)
	}

	w.WriteHeader(http.StatusCreated)
}

func (ctr *Controller) GetBatchEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodGet {
		http.Error(w, "wrong http method", http.StatusMethodNotAllowed)
		return
	}

	batch, err := ctr.GetBatch(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SuccessBatchRespond(batch, w)
}
