package error_handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Detail  string `json:"detail"`
}

const (
	DataCannotHandledError = "Data cannot handled from db"
	InvalidSchemaError     = "Schema is not correct"
)

func HandleDataCannotHandledError(w http.ResponseWriter, r *http.Request, log *log.Logger) {
	msg := DataCannotHandledError
	log.Println(msg)
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(&ErrorResponse{
		Message: msg,
		Status:  http.StatusInternalServerError,
	})
}

func HandleInvalidSchemaError(w http.ResponseWriter, r *http.Request, err error, log *log.Logger) {
	msg := InvalidSchemaError
	log.Println(msg)
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(&ErrorResponse{
		Message: msg,
		Detail:  err.Error(),
		Status:  http.StatusBadRequest,
	})
}
