package routes

import (
	"github.com/gorilla/mux"
	"mongoGo/internal/controller"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user", controller.CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/user/{id}", controller.GetUserEndpoint).Methods("GET")
	router.HandleFunc("/user", controller.GetUsersEndpoint).Methods("GET")
	//router.HandleFunc("/user/{id}", controller.DeleteUserEndpoint).Methods("DELETE")
	//router.HandleFunc("/user/{id}", controller.UpdateUserEndpoint).Methods("PUT")
	//router.HandleFunc("/upload", controller.UploadUserEndpoint).Methods("POST")
	return router
}
