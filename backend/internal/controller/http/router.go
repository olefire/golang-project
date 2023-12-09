package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (ctr *Controller) NewRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/user", ctr.CreateUserEndpoint).Methods("Post")
	router.HandleFunc("/user/{id}", ctr.GetUserEndpoint).Methods("Get")
	router.HandleFunc("/user/{id}", ctr.DeleteUserEndpoint).Methods("Delete")
	router.HandleFunc("/users", ctr.GetUsersEndpoint).Methods("Get")

	router.HandleFunc("/paste", ctr.CreatePasteEndpoint).Methods("Post")
	router.HandleFunc("/paste/{id}", ctr.GetPasteEndpoint).Methods("Get")
	router.HandleFunc("/paste/{id}", ctr.DeletePasteEndpoint).Methods("Delete")
	router.HandleFunc("/batch", ctr.GetBatchEndpoint).Methods("Get")

	router.HandleFunc("/lint", ctr.LintEndpoint).Methods("Post")

	return router
}
