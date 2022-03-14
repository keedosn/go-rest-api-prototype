package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/handler"
)

func TestCreateCategoryHandler(t *testing.T) {
	t.Run("test create category handler", func(t *testing.T) {
		body := "{\"name\":\"Test Category\"}"
		request := httptest.NewRequest(http.MethodPost, "/api/category", strings.NewReader(body))
		responseRecorder := httptest.NewRecorder()

		handler.CreateCategoryHandler.ServeHTTP(responseRecorder, request)

		if responseRecorder.Code != http.StatusCreated {
			t.Errorf("Want status '%d', got '%d'", http.StatusCreated, responseRecorder.Code)
		}
	})
}
