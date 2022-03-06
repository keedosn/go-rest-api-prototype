package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/handler"
)

func TestCreateArticleHandler(t *testing.T) {
	t.Run("test create article handler", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/api/article", nil)
		responseRecorder := httptest.NewRecorder()

		handler.CreateArticleHandler.ServeHTTP(responseRecorder, request)

		if responseRecorder.Code != http.StatusCreated {
			t.Errorf("Want status '%d', got '%d'", http.StatusCreated, responseRecorder.Code)
		}

		// if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
		// 	t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
		// }
	})
}
