package main

import (
	"log"
	"net/http"

	"github.com/AndresMaldoando24/go_htmx/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Route("/", func(r chi.Router) {
		r.Get("/", handlers.Home)
	})

	r.Route("/login", func(r chi.Router) {
		r.Get("/", handlers.Login)
		r.Post("/auth", handlers.Auth)
	})

	log.Println("App running on 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
