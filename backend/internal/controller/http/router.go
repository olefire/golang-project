package http

import (
	"backend/internal/middleware"
	"backend/internal/middleware/protected"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctr *Controller) NewRouter() http.Handler {
	router := gin.Default()
	router.POST("/auth/register", ctr.SignUpUserEndpoint)
	router.POST("/auth/login", ctr.SignInUserEndpoint)
	router.POST("/auth/logout", ctr.SignUpUserEndpoint)

	router.GET("/user/:id", ctr.GetUserEndpoint)
	router.DELETE("/user/:id", ctr.DeleteUserEndpoint)
	router.GET("/user", ctr.GetUsersEndpoint)
	router.GET("/me", middleware.DeserializeUser(), ctr.GetMe)

	router.POST("/paste", middleware.DeserializeUser(), ctr.CreatePasteEndpoint)
	router.GET("/paste/:id", middleware.DeserializeUser(), ctr.GetPasteEndpoint)

	router.PATCH("/paste/:id", middleware.DeserializeUser(), protected.ProtectedPaste(ctr.PasteService), ctr.UpdatePasteEndpoint)

	router.DELETE("/paste/:id", middleware.DeserializeUser(), protected.ProtectedPaste(ctr.PasteService), ctr.DeletePasteEndpoint)
	router.GET("/paste", ctr.GetBatchEndpoint)

	router.POST("/lint", ctr.LintEndpoint)

	return router
}
