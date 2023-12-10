package http

import (
	"backend/internal/models"
	"backend/internal/models/linter"
	"backend/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
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

func (ctr *Controller) SignUpUserEndpoint(ctx *gin.Context) {

	var user *models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	id, err := ctr.SignUpUser(ctx, user)

	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": "user successfully signed up", "id": id})
}

// Todo
//func (ctr *Controller) SignInUserEndpoint(ctx *gin.Context) {
//	var credentials *models.SignInInput
//
//	if err := ctx.ShouldBindJSON(&credentials); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
//		return
//	}
//
//	user, err := ctr.FindUserByEmail(ctx, credentials.Email)
//	if err != nil {
//		if errors.Is(err, mongo.ErrNoDocuments) {
//			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or password"})
//			return
//		}
//		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
//		return
//	}
//
//	if err := utils.VerifyPassword(user.Password, credentials.Password); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
//		return
//	}
//
//	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
//		return
//	}
//
//	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
//		return
//	}
//
//	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
//	ctx.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
//	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)
//
//	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
//}

func (ctr *Controller) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ctr *Controller) GetUsersEndpoint(ctx *gin.Context) {

	users, err := ctr.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": users})
}

func (ctr *Controller) GetUserEndpoint(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := ctr.GetUser(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

func (ctr *Controller) DeleteUserEndpoint(ctx *gin.Context) {
	id := ctx.Param("id")

	err := ctr.DeleteUser(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("deleted user id: %s", id), "id": id})
}

func (ctr *Controller) CreatePasteEndpoint(ctx *gin.Context) {
	var paste *models.Paste

	if err := ctx.ShouldBindJSON(&paste); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := ctr.CreatePaste(ctx, paste)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": id})
}

func (ctr *Controller) GetBatchEndpoint(ctx *gin.Context) {

	batch, err := ctr.GetBatch(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": batch})
}

func (ctr *Controller) GetPasteEndpoint(ctx *gin.Context) {
	id := ctx.Param("id")

	paste, err := ctr.GetPasteById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": paste})
}

func (ctr *Controller) DeletePasteEndpoint(ctx *gin.Context) {
	id := ctx.Param("id")

	err := ctr.DeletePaste(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("deleted user id: %s", id), "id": id})
}

func (ctr *Controller) LintEndpoint(ctx *gin.Context) {
	byteData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	sourceFile := linter.SourceFile{Code: string(byteData), Language: "python"}

	lintIssues, err := ctr.LintCode(sourceFile)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"LintResult": lintIssues})

}
