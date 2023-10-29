package http

import (
	"encoding/json"
	"mongoGo/internal/models"
	"net/http"
)

type errData[T any] struct {
	Data       T      `json:"data"`
	StatusCode int    `json:"status"`
	Message    string `json:"msg"`
}

func SuccessBatchRespond(batch []models.Paste, w http.ResponseWriter) {
	payload := errData[[]models.Paste]{
		Data:       batch,
		StatusCode: http.StatusOK,
		Message:    "success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func SuccessUsersRespond(users []models.User, w http.ResponseWriter) {
	payload := errData[[]models.User]{
		Data:       users,
		StatusCode: http.StatusOK,
		Message:    "success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func SuccessfulCreation(id string, w http.ResponseWriter) {
	payload := struct {
		Id string `json:"id"`
	}{Id: id}
	err := json.NewEncoder(w).Encode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusTeapot)
	}

	w.WriteHeader(http.StatusCreated)
}
