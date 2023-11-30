package api

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (a *Api) Cors() {
	a.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour),
	}))
}

func (a *Api) Routes() {
	a.router.Route("/email", func(r chi.Router) {
		r.Post("/add", a.mailController.Add)
	})
}

func (a *Api) GetRouter() *chi.Mux {
	return a.router
}
