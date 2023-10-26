package http

import (
	"net/http"
)

func (ctr *Controller) NewRouter() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/user", ctr.CreateUserEndpoint)
	router.HandleFunc("/users", ctr.GetUsersEndpoint)

	router.HandleFunc("/paste", ctr.CreatePasteEndpoint)
	router.HandleFunc("/batch", ctr.GetBatchEndpoint)

	return router
}
