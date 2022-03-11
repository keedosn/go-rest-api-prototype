package handler

import "net/http"

type NotFoundHandler struct{}

func (NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encodeResponse(w, &response{http.StatusNotFound, "Path " + r.RequestURI + " not found"}, nil)
}

type MethodNotAllowedHandler struct{}

func (MethodNotAllowedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encodeResponse(w, &response{http.StatusMethodNotAllowed, "Method Not Allowed: " + r.Method}, nil)
}
