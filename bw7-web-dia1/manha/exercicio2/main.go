package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Person struct {
	FistName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {
		var person Person

		if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		response := fmt.Sprintf("Hello, %v %v", person.FistName, person.LastName)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	})

	if err := http.ListenAndServe(":8081", r); err != nil {
		panic(err)
	}

}
