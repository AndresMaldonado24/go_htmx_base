package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	

	r.Route("/", func(r chi.Router){
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			tmpl := template.Must(template.ParseFiles("./templates/index.html"))
			tmpl.Execute(w, "Welcome to Go & HTMX")
		})
		
		r.Post("/clicked", func(w http.ResponseWriter, r *http.Request) {
			msg := []byte("You clicked!")
			w.Write(msg)
		})
	})

	log.Println("App running on 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}