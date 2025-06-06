package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(200)
		w.Write([]byte("pong"))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}

}
