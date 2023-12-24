package errors

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Detail  string `json:"detail"`
}
