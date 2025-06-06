package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var products []Product

func readFileJson(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error ao abrir o arquivo JSON: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error ao ler o arquivo JSON: %v", err)
	}

	err = json.Unmarshal(bytes, &products)
	if err != nil {
		log.Fatalf("Error ao fazer o parse do JSON: %v", err)
	}
	fmt.Println("Arquivo JSON lido com sucesso!", filename)
}

func main() {
	readFileJson("bw7-web-dia1/tarde/products.json")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	r.Get("/products", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
		json.NewEncoder(w).Encode(products)
	})

	r.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {

		paramId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		for _, product := range products {
			if product.ID == id {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(product)
				return
			}
		}
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
	})

	r.Get("/products/search", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
	fmt.Println("Servidor rodando na porta 8080")
}
