package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctr *Controller) NewRouter() http.Handler {
	//router := mux.NewRouter()
	router := gin.Default()
	router.POST("/auth/register", ctr.SignUpUserEndpoint)
	//router.POST("/auth/login", ctr.SignInUserEndpoint)
	router.POST("/auth/logout", ctr.SignUpUserEndpoint)

	router.GET("/user/{id}", ctr.GetUserEndpoint)
	router.DELETE("/user/{id}", ctr.DeleteUserEndpoint)
	router.GET("/user", ctr.GetUsersEndpoint)

	router.POST("/paste", ctr.CreatePasteEndpoint)
	router.GET("/paste/{id}", ctr.GetPasteEndpoint)
	router.DELETE("/paste/{id}", ctr.DeletePasteEndpoint)
	router.GET("/paste", ctr.GetBatchEndpoint)

	router.POST("/lint", ctr.LintEndpoint)

	return router
}
