package main

import (
	"github.com/rs/cors"
	"log"
	"mongoGo/internal/routes"
	"mongoGo/pkg/middleware"
	"net/http"
)

func main() {

	router := routes.Routes()

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	handler := c.Handler(router)
	err := http.ListenAndServe(":8080", middleware.LogRequest(handler))
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}

}
