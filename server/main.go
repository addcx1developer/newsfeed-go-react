package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	r := chi.NewRouter()

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: false,
	})

	r.Get("/graphql", h.ServeHTTP)
	r.Post("/graphql", h.ServeHTTP)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
