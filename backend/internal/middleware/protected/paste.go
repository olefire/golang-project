package protected

import (
	"backend/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProtectedPaste(pasteService services.PasteManagement) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		currentUserId := ctx.MustGet("currentUser")
		curId := fmt.Sprint(currentUserId)

		paste, err := pasteService.GetPasteById(ctx, id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		if paste.UserID.Hex() != curId {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "you don't have permission to access this resource"})
			err = fmt.Errorf("you don't have permission to access this resource")
		}

		if err != nil {
			ctx.Abort()
		}

		ctx.Next()
	}
}
