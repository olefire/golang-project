package middleware

import (
	"backend/internal/config"
	utils "backend/internal/utils/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// DeserializeUser todo remove config
func DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = cookie
		}

		if accessToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		cfg := config.NewConfig()
		sub, err := utils.ValidateToken(accessToken, cfg.AccessTokenPublicKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.Set("currentUser", fmt.Sprint(sub))
		ctx.Next()
	}
}
