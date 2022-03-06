package handler

import "net/http"

type ErrorResponse struct {
	Error string `json:"error"`
}

func Error(err string) *ErrorResponse {
	return &ErrorResponse{Error: err}
}

type NotFoundHandler struct{}

func (NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encodeResponse(w, &response{http.StatusNotFound, "Path " + r.RequestURI + " not found"}, nil)
}

type MethodNotAllowedHandler struct{}

func (MethodNotAllowedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encodeResponse(w, &response{http.StatusMethodNotAllowed, "Method Not Allowed: " + r.Method}, nil)
}
