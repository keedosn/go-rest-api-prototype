package handler

import "net/http"

type ErrorResponse struct {
	Message string `json:"message"`
}

type NotFoundHandler struct{}

func (NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	encodeResponse(w, &ErrorResponse{Message: "Path " + r.RequestURI + " not found"}, nil)
}

type MethodNotAllowedHandler struct{}

func (MethodNotAllowedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	encodeResponse(w, &ErrorResponse{Message: "Method Not Allowed: " + r.Method}, nil)
}
