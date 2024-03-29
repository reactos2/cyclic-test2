package main

import (
	"net/http"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const dsn = "postgresql://3d091b03-47ef-4aca-b9ea-b5869d0cde02-user:pw-b1686ed8-8893-4fed-b760-189d69e246a5@postgres-free-tier-v2020.gigalixir.com:5432/3d091b03-47ef-4aca-b9ea-b5869d0cde02"

func main() {
	err := modDatabase.CDBInitialize(dsn)
	if err != nil {
		fmt.Println(err)
		return
	} else {
    fmt.Println("db init successful")
	}
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