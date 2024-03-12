package http

import (
	"backend/internal/config"
	"backend/internal/models"
	"backend/internal/models/linter"
	"backend/internal/services"
	utils "backend/internal/utils/auth"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"net/http"
	"strings"
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

// SignInUserEndpoint todo remove config
func (ctr *Controller) SignInUserEndpoint(ctx *gin.Context) {
	var credentials *models.SignInInput

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	user, err := ctr.FindUserByEmail(ctx, credentials.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or password"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := utils.VerifyPassword(user.Password, credentials.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	cfg := config.NewConfig()

	accessToken, err := utils.CreateToken(time.Hour*24*7, user.ID, cfg.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refreshToken, err := utils.CreateToken(time.Hour*24*7*30, user.ID, cfg.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("accessToken", accessToken, 30*24*60*60, "/", "localhost", false, true)
	ctx.SetCookie("refreshToken", refreshToken, 30*24*60*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", 30*24*60*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "accessToken": accessToken})
}

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
	currentUser := ctx.MustGet("currentUser")

	var paste *models.Paste

	if err := ctx.ShouldBindJSON(&paste); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	currentId, err := primitive.ObjectIDFromHex(fmt.Sprint(currentUser))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	paste.UserID = currentId

	id, err := ctr.CreatePaste(ctx, paste)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": paste, "id": currentId})
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

	currentUserId := ctx.MustGet("currentUser")
	curId := fmt.Sprint(currentUserId)

	paste, err := ctr.GetPasteById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if paste.UserID.Hex() != curId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "you don't have permission to access this resource"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": paste})
}

func (ctr *Controller) DeletePasteEndpoint(ctx *gin.Context) {
	id := ctx.Param("id")

	currentUserId := ctx.MustGet("currentUser")
	curId := fmt.Sprint(currentUserId)

	paste, err := ctr.GetPasteById(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if paste.UserID.Hex() != curId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "you don't have permission to access this resource"})
		return
	}

	err = ctr.DeletePaste(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("deleted user id: %s", id), "id": id})
}

func (ctr *Controller) UpdatePasteEndpoint(ctx *gin.Context) {
	id := ctx.Param("id")

	var updPaste *models.UpdatePaste

	if err := ctx.ShouldBindJSON(&updPaste); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	paste, err := ctr.UpdatePaste(ctx, id, updPaste)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "updated paste": paste})
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

func (ctr *Controller) GetMe(ctx *gin.Context) {
	currentUserId := ctx.MustGet("currentUser")

	user, err := ctr.GetUser(ctx, fmt.Sprint(currentUserId))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": user}})
}
