package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"

	"github.com/addcx1developer/newsfeed-go-react/server/data"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
)

func main() {
	result := graphql.Do(graphql.Params{
		Schema:        data.Schema,
		RequestString: testutil.IntrospectionQuery,
	})

	if len(result.Errors) > 0 {
		log.Fatalf("Introspection error: %+v", result.Errors)
	}

	jsonFile := "./server/data/schema.json"
	file, err := os.Create(jsonFile)

	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}

	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	if err := enc.Encode(result.Data); err != nil {
		log.Fatalf("Failed to encode JSON: %v", err)
	}

	log.Println("Saved introspection JSON to", jsonFile)

	cmd := exec.Command("node", "./scripts/printSDL.mjs", jsonFile, "./server/data/schema.graphql")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to generate SDL: %v", err)
	}

	log.Println("Generated SDL at ./server/data/schema.graphql")
}
