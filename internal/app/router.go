package app

import (
	"net/http"

	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/handler"
	"github.com/gorilla/mux"
)

func SetupRouter(env *handler.Env) *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = &handler.NotFoundHandler{}
	r.MethodNotAllowedHandler = &handler.MethodNotAllowedHandler{}

	r.Use(PrepareHeadersMiddleware)
	r.Use(ValidateJsonBodyMiddleware) // probably not needed
	r.Use(LoggingMiddleware)

	hc := r.PathPrefix("/health").Subrouter()
	hc.Handle("", handler.New(env, handler.HealthCheckHandler)).Methods(http.MethodGet)

	api := r.PathPrefix("/api").Subrouter()
	api.Handle("/article", handler.New(env, handler.CreateArticleHandler)).Methods(http.MethodPost)

	api.Handle("/category", handler.New(env, handler.CreateCategoryHandler)).Methods(http.MethodPost)

	return r
}
