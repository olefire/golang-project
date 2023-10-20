package middleware

import (
	"encoding/json"
	"mongoGo/internal/models"
	"net/http"
)

type errData struct {
	StatusCode int    `json:"status"`
	Message    string `json:"msg"`
}

// AuthorizationResponse -> response authorize
func AuthorizationResponse(msg string, writer http.ResponseWriter) {
	temp := &errData{StatusCode: http.StatusUnauthorized, Message: msg}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)
	err := json.NewEncoder(writer).Encode(temp)
	if err != nil {
		return
	}
}

// SuccessArrRespond -> response formatter
func SuccessArrRespond(fields []*models.User, writer http.ResponseWriter) {
	// var fields["status"] := "success"
	_, err := json.Marshal(fields)
	if err != nil {
		return
	}
	type data struct {
		People     []*models.User `json:"data"`
		StatusCode int            `json:"status"`
		Message    string         `json:"msg"`
	}
	temp := &data{People: fields, StatusCode: http.StatusOK, Message: "success"}
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
		return
	}
}
