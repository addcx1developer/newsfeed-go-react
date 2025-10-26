package main

import (
	"encoding/json"
	"log"
	"os"

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
		log.Fatalf("Failed to introspect schema: %v", result.Errors)
	}

	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling schema: %v", err)
	}

	err = os.WriteFile("./server/data/schema.json", b, 0644)
	if err != nil {
		log.Fatalf("Error writing schema.json: %v", err)
	}

	log.Println("✅ schema.json generated successfully!")
}
