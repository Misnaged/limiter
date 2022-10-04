package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"limiter/config"
	"limiter/internal/service"
	"net/http"
)

func NewServer(config *config.Scheme, srv *service.Service) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: Router(srv),
	}
}

func Router(srv *service.Service) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/", func(r chi.Router) {
		r.Get("/signup", srv.NewCredentials())
		go http.ListenAndServe(":3000", r)
	})
	return r
}
