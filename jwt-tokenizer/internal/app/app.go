package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/internal/endpoint"
	"github.com/idkwhyureadthis/food-delivery/jwt-tokenizer/internal/service"
)

type App struct {
	r *chi.Mux
	e *endpoint.Endpoint
	s *service.Service
}

func New(dbPath string) (*App, error) {
	a := &App{}
	a.s = service.New(dbPath)
	a.e = endpoint.New(a.s)
	a.r = chi.NewRouter()
	a.r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	return a, nil
}

func (a *App) Run(port string) error {
	log.Printf("server running at http://localhost%s", port)
	srv := &http.Server{
		Addr:    port,
		Handler: a.r,
	}
	return srv.ListenAndServe()
}
