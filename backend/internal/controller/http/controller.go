package http

import (
	"backend/internal/models"
	"backend/internal/models/linter"
	"backend/internal/services"
	utils "backend/internal/utils/auth"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"time"
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

type AuthService struct {
	services.AuthManagement
}

type Controller struct {
	UserService
	PasteService
	LinterService
	AuthService
}

func NewController(us UserService, ps PasteService, ls LinterService, as AuthService) *Controller {
	return &Controller{
		UserService:   us,
		PasteService:  ps,
		LinterService: ls,
		AuthService:   as,
	}
}

func (ctr *Controller) SignUpUserEndpoint(w http.ResponseWriter, r *http.Request) {
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

	id, err := ctr.SignUpUser(ctx, &req)
	if err != nil {
		http.Error(w, "can`t sign up user", http.StatusInternalServerError)
		return
	}

	SuccessfulCreation(id, w)
}

func (ctr *Controller) SignInUserEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req *models.SignInInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := ctr.FindUserByEmail(ctx, req.Email)
	if err != nil {
		//if errors.Is(err, mongo.ErrNoDocuments) {
		//	http.Error(w, "no such email email", http.StatusBadRequest)
		//	return
		//}
		http.Error(w, "no such email "+req.Email, http.StatusBadRequest)
		return
	}

	if err := utils.VerifyPassword(user.Password, req.Password); err != nil {
		http.Error(w, "invalid email or password", http.StatusBadRequest)
		return
	}

	//config, _ := config.LoadConfig(".")
	// Generate Tokens
	accessToken, err := utils.CreateToken(time.Duration(time.Minute*15), user.ID, "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlCUEFJQkFBSkJBTzVIKytVM0xrWC91SlRvRHhWN01CUURXSTdGU0l0VXNjbGFFKzlaUUg5Q2VpOGIxcUVmCnJxR0hSVDVWUis4c3UxVWtCUVpZTER3MnN3RTVWbjg5c0ZVQ0F3RUFBUUpCQUw4ZjRBMUlDSWEvQ2ZmdWR3TGMKNzRCdCtwOXg0TEZaZXMwdHdtV3Vha3hub3NaV0w4eVpSTUJpRmI4a25VL0hwb3piTnNxMmN1ZU9wKzVWdGRXNApiTlVDSVFENm9JdWxqcHdrZTFGY1VPaldnaXRQSjNnbFBma3NHVFBhdFYwYnJJVVI5d0loQVBOanJ1enB4ckhsCkUxRmJxeGtUNFZ5bWhCOU1HazU0Wk1jWnVjSmZOcjBUQWlFQWhML3UxOVZPdlVBWVd6Wjc3Y3JxMTdWSFBTcXoKUlhsZjd2TnJpdEg1ZGdjQ0lRRHR5QmFPdUxuNDlIOFIvZ2ZEZ1V1cjg3YWl5UHZ1YStxeEpXMzQrb0tFNXdJZwpQbG1KYXZsbW9jUG4rTkVRdGhLcTZuZFVYRGpXTTlTbktQQTVlUDZSUEs0PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	refreshToken, err := utils.CreateToken(time.Duration(time.Minute*60), user.ID, "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlCT1FJQkFBSkJBSWFJcXZXeldCSndnYjR1SEhFQ01RdHFZMTI5b2F5RzVZMGlGcG51a0J1VHpRZVlQWkE4Cmx4OC9lTUh3Rys1MlJGR3VxMmE2N084d2s3TDR5dnY5dVY4Q0F3RUFBUUpBRUZ6aEJqOUk3LzAxR285N01CZUgKSlk5TUJLUEMzVHdQQVdwcSswL3p3UmE2ZkZtbXQ5NXNrN21qT3czRzNEZ3M5T2RTeWdsbTlVdndNWXh6SXFERAplUUloQVA5UStrMTBQbGxNd2ZJbDZtdjdTMFRYOGJDUlRaZVI1ZFZZb3FTeW40YmpBaUVBaHVUa2JtZ1NobFlZCnRyclNWZjN0QWZJcWNVUjZ3aDdMOXR5MVlvalZVRlVDSUhzOENlVHkwOWxrbkVTV0dvV09ZUEZVemhyc3Q2Z08KU3dKa2F2VFdKdndEQWlBdWhnVU8yeEFBaXZNdEdwUHVtb3hDam8zNjBMNXg4d012bWdGcEFYNW9uUUlnQzEvSwpNWG1heWtsaFRDeWtXRnpHMHBMWVdkNGRGdTI5M1M2ZUxJUlNIS009Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})
	SuccessLogin(accessToken, w)
}

func (ctr *Controller) LogoutUserEndpoint(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})

	//ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	//ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	//ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, true)
	//
	//ctx.JSON(http.StatusOK, gin.H{"status": "success"})
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

	byteData, err := io.ReadAll(r.Body)
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
