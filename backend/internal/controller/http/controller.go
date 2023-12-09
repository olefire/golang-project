package http

import (
	"backend/internal/models"
	"backend/internal/models/linter"
	"backend/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type UserService struct {
	services.UserManagement
}

type PasteService struct {
	services.PasteManagement
}

type LinterService struct {
	services.Linter
}

type Controller struct {
	UserService
	PasteService
	LinterService
}

func NewController(us UserService, ps PasteService, ls LinterService) *Controller {
	return &Controller{
		UserService:   us,
		PasteService:  ps,
		LinterService: ls,
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

	SuccessfulCreation(id, w)
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

	SuccessUsersRespond(users, w)
}

func (ctr *Controller) GetUserEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "wrong http method", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := ctr.GetUser(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SuccessUserRespond(user, w)
}

func (ctr *Controller) DeleteUserEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "wrong http method", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	err := ctr.DeleteUser(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	SuccessDelete(id, w)
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

	SuccessfulCreation(id, w)
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

func (ctr *Controller) GetPasteEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "wrong http method", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	paste, err := ctr.GetPasteById(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SuccessPasteRespond(paste, w)
}

func (ctr *Controller) DeletePasteEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "wrong http method", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	err := ctr.DeletePaste(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SuccessDelete(id, w)
}

func (ctr *Controller) LintEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "wrong http method", http.StatusMethodNotAllowed)
		return
	}

	byteData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sourceFile := linter.SourceFile{Code: string(byteData), Language: "python"}

	lintIssues, err := ctr.LintCode(sourceFile)

	err = json.NewEncoder(w).Encode(&lintIssues)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
