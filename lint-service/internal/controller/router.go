package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctr *Controller) NewRouter() http.Handler {
	router := gin.Default()

	router.POST("/lint", ctr.LintEndpoint)

	return router
}
