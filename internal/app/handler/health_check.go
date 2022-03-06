package handler

import (
	"net/http"
)

var HealthCheckHandler *Handler

func init() {
	HealthCheckHandler = &Handler{
		Handle:   HealthCheckHandlerFunc,
		Request:  &HealthCheckRequest{},
		Response: &HealthCheckResponse{},
	}
}

type HealthCheckRequest struct {
}

type HealthCheckResponse struct {
	Status string                   `json:"status"`
	Data   *HealthCheckResponseBody `json:"data"`
}

type HealthCheckResponseBody struct {
	Message string `json:"message,omitempty"`
}

func HealthCheckHandlerFunc(_ *Handler, w http.ResponseWriter) (interface{}, int, error) {
	return &HealthCheckResponseBody{
		Message: "This is welcome health message. Everything seems to be alright ;)",
	}, http.StatusOK, nil

	// return &HealthCheckResponse{
	// 	Status: http.StatusText(http.StatusOK),
	// 	Data: &HealthCheckResponseBody{
	// 		Message: "This is welcome health message. Everything seems to be alright ;)",
	// 	},
	// }, http.StatusOK, nil
}
