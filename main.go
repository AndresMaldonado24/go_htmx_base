package main

import (
	"log"
	"net/http"

	hl "github.com/AndresMaldoando24/go_htmx_base/handlers"
	h "github.com/AndresMaldoando24/go_htmx_base/helpers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.With(h.CheckCookie).Route("/", func(r chi.Router) {
		r.Get("/", hl.Home)
	})

	r.With(h.CheckCookie).Route("/login", func(r chi.Router) {
		r.Get("/", hl.Login)
		r.Post("/", hl.Auth)
	})

	log.Println("App running on 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
