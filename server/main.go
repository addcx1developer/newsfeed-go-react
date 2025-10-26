package main

import (
	"log"
	"net/http"

	"github.com/addcx1developer/newsfeed-go-react/server/data"
	"github.com/go-chi/chi/v5"
	"github.com/graphql-go/handler"
)

func main() {
	r := chi.NewRouter()

	h := handler.New(&handler.Config{
		Schema:     &data.Schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: false,
	})

	r.Get("/graphql", h.ServeHTTP)
	r.Post("/graphql", h.ServeHTTP)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
