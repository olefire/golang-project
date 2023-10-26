package http

import (
	"encoding/json"
	"mongoGo/internal/models"
	"net/http"
)

type errData struct {
	StatusCode int    `json:"status"`
	Message    string `json:"msg"`
}

func SuccessBatchRespond(batch []models.Paste, w http.ResponseWriter) {
	_, err := json.Marshal(batch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	type data struct {
		Batch      []models.Paste `json:"data"`
		StatusCode int            `json:"status"`
		Message    string         `json:"msg"`
	}
	temp := &data{Batch: batch, StatusCode: http.StatusOK, Message: "success"}
	if err != nil {
		ServerErrResponse(err.Error(), w)
	}

	//Send header, status code and output to writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(temp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

// SuccessArrRespond -> response formatter
func SuccessArrRespond(fields []models.User, writer http.ResponseWriter) {
	// var fields["status"] := "success"
	_, err := json.Marshal(fields)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	type data struct {
		People     []models.User `json:"data"`
		StatusCode int           `json:"status"`
		Message    string        `json:"msg"`
	}
	temp := &data{People: fields, StatusCode: http.StatusOK, Message: "success"}
	if err != nil {
		ServerErrResponse(err.Error(), writer)
		return
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(temp)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

// SuccessRespond -> response formatter
func SuccessRespond(fields models.User, writer http.ResponseWriter) {
	_, err := json.Marshal(fields)
	type data struct {
		Person     models.User `json:"data"`
		StatusCode int         `json:"status"`
		Message    string      `json:"msg"`
	}
	temp := &data{Person: fields, StatusCode: http.StatusOK, Message: "success"}
	if err != nil {
		ServerErrResponse(err.Error(), writer)
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode(temp)
	if err != nil {
		return
	}
}

// SuccessResponse -> success formatter
func SuccessResponse(msg string, writer http.ResponseWriter) {
	temp := &errData{StatusCode: http.StatusOK, Message: msg}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(temp)
	if err != nil {
		return
	}
}

// ErrorResponse -> error formatter
func ErrorResponse(error string, writer http.ResponseWriter) {
	temp := &errData{StatusCode: http.StatusBadRequest, Message: error}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	err := json.NewEncoder(writer).Encode(temp)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

// ServerErrResponse -> server error formatter
func ServerErrResponse(error string, writer http.ResponseWriter) {
	type serverErrData struct {
		StatusCode int    `json:"status"`
		Message    string `json:"msg"`
	}
	temp := &serverErrData{StatusCode: http.StatusInternalServerError, Message: error}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(writer).Encode(temp)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

// ValidationResponse -> user input validation
func ValidationResponse(fields map[string][]string, writer http.ResponseWriter) {
	//Create a new map and fill it
	response := make(map[string]interface{})
	response["errors"] = fields
	response["status"] = http.StatusUnprocessableEntity
	response["msg"] = "validation error"

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnprocessableEntity)
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}
