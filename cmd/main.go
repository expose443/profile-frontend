package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/with-insomnia/profile-frontend/internal/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	h := handlers.NewHandler()
	r.Get("/login", h.LoginGet)
	r.Post("/login", h.LoginPost)
	fs := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fs))

	log.Fatal(http.ListenAndServe(":3000", r))
}
