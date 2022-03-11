package definition

type ErrorResponse struct {
	Error string `json:"error"`
}

func Error(err string) *ErrorResponse {
	return &ErrorResponse{err}
}
