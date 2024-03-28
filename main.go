package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
  r.Get("/test2", handleHello2)
  
	http.ListenAndServe(":3000", r)
}

func handleHello2(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("hello2"))
}