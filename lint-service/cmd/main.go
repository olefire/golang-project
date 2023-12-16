package main

import (
	"github.com/rs/cors"
	"lint-service/internal/controller"
	"lint-service/internal/linters/python/metrics"
	"lint-service/internal/linters/python/pylint"
	"lint-service/internal/services"
	"lint-service/internal/services/linter"
	"log"
	"net/http"
)

func main() {
	pyLint := pylint.Linter{}
	pyMetrics := metrics.Linter{}

	linters := []services.Linter{&pyLint, &pyMetrics}

	linterService := linter.NewClient(linters)

	ctr := controller.NewController(controller.LinterService{Service: linterService})

	router := ctr.NewRouter()

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	handler := c.Handler(router)

	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
