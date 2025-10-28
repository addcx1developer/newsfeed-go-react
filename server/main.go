package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/addcx1developer/newsfeed-go-react/server/data"
	"github.com/go-chi/chi/v5"
	"github.com/graphql-go/handler"
)

var (
	persistedQueries map[string]string
	mu               sync.RWMutex
)

func loadPersistedQueries(path string) {
	dataBytes, err := os.ReadFile(path)
	if err != nil {
		log.Println("Failed to load persisted queries:", err)
		return
	}

	var queries map[string]string
	if err := json.Unmarshal(dataBytes, &queries); err != nil {
		log.Println("Invalid persisted queries JSON:", err)
		return
	}

	mu.Lock()
	persistedQueries = queries
	mu.Unlock()

	log.Printf("Loaded %d persisted queries\n", len(queries))
}

func persistedHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var body map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if docID, ok := body["doc_id"].(string); ok && body["query"] == nil {
				mu.RLock()
				query, found := persistedQueries[docID]
				mu.RUnlock()
				if found {
					body["query"] = query
				} else {
					http.Error(w, "Persisted query not found", http.StatusNotFound)
					return
				}
			}

			newBody, _ := json.Marshal(body)
			r.Body = io.NopCloser(bytes.NewReader(newBody))
			r.ContentLength = int64(len(newBody))
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	const queriesPath = "./persisted-queries.json"

	loadPersistedQueries(queriesPath)

	h := handler.New(&handler.Config{
		Schema:     &data.Schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: false,
	})

	r := chi.NewRouter()

	r.Post("/graphql", func(w http.ResponseWriter, r *http.Request) {
		persistedHandler(h).ServeHTTP(w, r)
	})

	r.Get("/graphql", func(w http.ResponseWriter, r *http.Request) {
		persistedHandler(h).ServeHTTP(w, r)
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
