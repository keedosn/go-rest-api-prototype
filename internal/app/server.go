package app

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	def "git.pbiernat.dev/golang/rest-api-prototype/internal/app/definition"
	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/handler"
)

type Server struct {
	*http.Server
}

func NewServer(env *handler.Env) *Server {
	return &Server{
		&http.Server{
			Handler:      SetupRouter(env),
			Addr:         env.Addr,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (s *Server) Start() {
	log.Println("Server listening on " + s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func PrepareHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Keep-Alive", "timeout=5")

		next.ServeHTTP(w, r)
	})
}

func ValidateJsonBodyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf, _ := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewReader(buf)) // rollack *Request to original state

		if len(buf) > 0 && !json.Valid(buf) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(def.Error("Unable to parse JSON: " + string(buf)))

			return
		}

		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request: " + r.RequestURI + " remote: " + r.RemoteAddr + " via: " + r.UserAgent())
		next.ServeHTTP(w, r)
	})
}
