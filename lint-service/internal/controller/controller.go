package controller

import (
	"github.com/gin-gonic/gin"
	"io"
	"lint-service/internal/models"
	"lint-service/internal/services/linter"
	"net/http"
)

type LinterService struct {
	*linter.Service
}

type Controller struct {
	LinterService
}

func NewController(ls LinterService) *Controller {
	return &Controller{
		LinterService: ls,
	}
}

func (ctr *Controller) LintEndpoint(ctx *gin.Context) {
	byteData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	sourceFile := models.SourceFile{Code: string(byteData), Language: "python"}

	lintIssues, err := ctr.LintCode(sourceFile)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, lintIssues)

}
