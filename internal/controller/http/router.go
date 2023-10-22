package http

import (
	"net/http"
)

func (ctr *Controller) NewRouter() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/user", ctr.CreateUserEndpoint)
	router.HandleFunc("/users", ctr.GetUsersEndpoint)

	return router
}
