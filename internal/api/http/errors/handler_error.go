package errors

import (
	"encoding/json"
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

func HandleDataCannotHandledError(w http.ResponseWriter, r *http.Request) string {
	msg := DataCannotHandledError
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(&ErrorResponse{
		Message: msg,
		Status:  http.StatusInternalServerError,
	})
	return msg
}

func HandleInvalidSchemaError(w http.ResponseWriter, r *http.Request, err error) string {
	msg := InvalidSchemaError
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(&ErrorResponse{
		Message: msg,
		Detail:  err.Error(),
		Status:  http.StatusBadRequest,
	})
	return msg
}
