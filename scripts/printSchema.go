package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/addcx1developer/newsfeed-go-react/server/data"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
)

func introspectionToSDL(data map[string]interface{}) string {
	var buf bytes.Buffer

	types := data["__schema"].(map[string]interface{})["types"].([]interface{})
	sort.Slice(types, func(i, j int) bool {
		return types[i].(map[string]interface{})["name"].(string) < types[j].(map[string]interface{})["name"].(string)
	})

	for _, t := range types {
		typ := t.(map[string]interface{})
		name := typ["name"].(string)
		kind := typ["kind"].(string)

		if strings.HasPrefix(name, "__") {
			continue
		}

		switch kind {
		case "OBJECT":
			buf.WriteString(fmt.Sprintf("type %s {\n", name))
			fields, ok := typ["fields"].([]interface{})
			if ok {
				for _, f := range fields {
					field := f.(map[string]interface{})
					fieldName := field["name"].(string)
					fieldType := unwrapType(field["type"].(map[string]interface{}))
					buf.WriteString(fmt.Sprintf("  %s: %s\n", fieldName, fieldType))
				}
			}
			buf.WriteString("}\n\n")
		case "ENUM":
			buf.WriteString(fmt.Sprintf("enum %s {\n", name))
			values, ok := typ["enumValues"].([]interface{})
			if ok {
				for _, v := range values {
					val := v.(map[string]interface{})
					buf.WriteString(fmt.Sprintf("  %s\n", val["name"].(string)))
				}
			}
			buf.WriteString("}\n\n")
		case "INTERFACE":
			buf.WriteString(fmt.Sprintf("interface %s {\n", name))
			fields, ok := typ["fields"].([]interface{})
			if ok {
				for _, f := range fields {
					field := f.(map[string]interface{})
					fieldName := field["name"].(string)
					fieldType := unwrapType(field["type"].(map[string]interface{}))
					buf.WriteString(fmt.Sprintf("  %s: %s\n", fieldName, fieldType))
				}
			}
			buf.WriteString("}\n\n")
		case "UNION":
			buf.WriteString(fmt.Sprintf("union %s = ", name))
			possibleTypes, ok := typ["possibleTypes"].([]interface{})
			if ok {
				names := []string{}
				for _, pt := range possibleTypes {
					names = append(names, pt.(map[string]interface{})["name"].(string))
				}
				buf.WriteString(strings.Join(names, " | "))
			}
			buf.WriteString("\n\n")
		case "SCALAR":
			buf.WriteString(fmt.Sprintf("scalar %s\n\n", name))
		case "INPUT_OBJECT":
			buf.WriteString(fmt.Sprintf("input %s {\n", name))
			inputFields, ok := typ["inputFields"].([]interface{})
			if ok {
				for _, f := range inputFields {
					field := f.(map[string]interface{})
					fieldName := field["name"].(string)
					fieldType := unwrapType(field["type"].(map[string]interface{}))
					buf.WriteString(fmt.Sprintf("  %s: %s\n", fieldName, fieldType))
				}
			}
			buf.WriteString("}\n\n")
		}
	}

	return buf.String()
}

func unwrapType(t map[string]interface{}) string {
	kind := t["kind"].(string)
	switch kind {
	case "NON_NULL":
		return fmt.Sprintf("%s!", unwrapType(t["ofType"].(map[string]interface{})))
	case "LIST":
		return fmt.Sprintf("[%s]", unwrapType(t["ofType"].(map[string]interface{})))
	default:
		return t["name"].(string)
	}
}

func main() {
	result := graphql.Do(graphql.Params{
		Schema:        data.Schema,
		RequestString: testutil.IntrospectionQuery,
	})

	if len(result.Errors) > 0 {
		log.Fatalf("Failed to introspect schema: %v", result.Errors)
	}

	jsonBytes, err := json.Marshal(result.Data)
	if err != nil {
		log.Fatalf("marshal error: %v", err)
	}

	var introspection map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &introspection); err != nil {
		log.Fatalf("unmarshal error: %v", err)
	}

	sdl := introspectionToSDL(introspection)
	if err := os.WriteFile("./server/data/schema.graphql", []byte(sdl), 0644); err != nil {
		log.Fatalf("write error: %v", err)
	}

	log.Println("✅ schema.graphql generated successfully (Go-only)!")
}
