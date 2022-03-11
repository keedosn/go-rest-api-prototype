package app

import (
	"net/http"

	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/handler"
	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/service"
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

	auth := r.PathPrefix("/auth").Subrouter()
	auth.Handle("/login", handler.New(env, handler.AuthLoginHandler)).Methods(http.MethodPost)

	api := r.PathPrefix("/api").Subrouter()
	api.Use(service.AuthService.ValidateUserTokenMiddleware) // only /api/** endpoints use this middleware

	api.Handle("/article", handler.New(env, handler.CreateArticleHandler)).Methods(http.MethodPost)

	api.Handle("/category", handler.New(env, handler.CreateCategoryHandler)).Methods(http.MethodPost)
	api.Handle("/category/{id:[0-9]+}", handler.New(env, handler.DeleteCategoryHandler)).Methods(http.MethodDelete)

	return r
}
