package main

import (
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

func unwrapType(t map[string]interface{}) string {
	if t == nil {
		return "Unknown"
	}
	switch t["kind"] {
	case "NON_NULL":
		return fmt.Sprintf("%s!", unwrapType(t["ofType"].(map[string]interface{})))
	case "LIST":
		return fmt.Sprintf("[%s]", unwrapType(t["ofType"].(map[string]interface{})))
	default:
		if name, ok := t["name"].(string); ok {
			return name
		}
		return "Unknown"
	}
}

func convertIntrospectionToSDL(introspection map[string]interface{}) string {
	var sb strings.Builder
	types := introspection["__schema"].(map[string]interface{})["types"].([]interface{})

	sort.Slice(types, func(i, j int) bool {
		return types[i].(map[string]interface{})["name"].(string) <
			types[j].(map[string]interface{})["name"].(string)
	})

	for _, t := range types {
		typeObj := t.(map[string]interface{})
		name := typeObj["name"].(string)
		kind := typeObj["kind"].(string)

		if strings.HasPrefix(name, "__") {
			continue
		}

		switch kind {
		case "OBJECT":
			sb.WriteString(fmt.Sprintf("type %s", name))
			if interfaces, ok := typeObj["interfaces"].([]interface{}); ok && len(interfaces) > 0 {
				names := []string{}
				for _, iface := range interfaces {
					names = append(names, iface.(map[string]interface{})["name"].(string))
				}
				sb.WriteString(" implements " + strings.Join(names, " & "))
			}
			sb.WriteString(" {\n")
			if fields, ok := typeObj["fields"].([]interface{}); ok {
				for _, f := range fields {
					field := f.(map[string]interface{})
					fieldName := field["name"].(string)
					fieldType := unwrapType(field["type"].(map[string]interface{}))
					sb.WriteString(fmt.Sprintf("  %s: %s\n", fieldName, fieldType))
				}
			}
			sb.WriteString("}\n\n")

		case "INTERFACE":
			sb.WriteString(fmt.Sprintf("interface %s {\n", name))
			if fields, ok := typeObj["fields"].([]interface{}); ok {
				for _, f := range fields {
					field := f.(map[string]interface{})
					sb.WriteString(fmt.Sprintf("  %s: %s\n",
						field["name"].(string),
						unwrapType(field["type"].(map[string]interface{})),
					))
				}
			}
			sb.WriteString("}\n\n")

		case "ENUM":
			sb.WriteString(fmt.Sprintf("enum %s {\n", name))
			if values, ok := typeObj["enumValues"].([]interface{}); ok {
				for _, v := range values {
					val := v.(map[string]interface{})
					sb.WriteString(fmt.Sprintf("  %s\n", val["name"].(string)))
				}
			}
			sb.WriteString("}\n\n")

		case "UNION":
			sb.WriteString(fmt.Sprintf("union %s = ", name))
			if possible, ok := typeObj["possibleTypes"].([]interface{}); ok {
				names := []string{}
				for _, p := range possible {
					names = append(names, p.(map[string]interface{})["name"].(string))
				}
				sb.WriteString(strings.Join(names, " | "))
			}
			sb.WriteString("\n\n")

		case "INPUT_OBJECT":
			sb.WriteString(fmt.Sprintf("input %s {\n", name))
			if fields, ok := typeObj["inputFields"].([]interface{}); ok {
				for _, f := range fields {
					field := f.(map[string]interface{})
					sb.WriteString(fmt.Sprintf("  %s: %s\n",
						field["name"].(string),
						unwrapType(field["type"].(map[string]interface{})),
					))
				}
			}
			sb.WriteString("}\n\n")

		case "SCALAR":
			sb.WriteString(fmt.Sprintf("scalar %s\n\n", name))
		}
	}

	return sb.String()
}

func main() {
	result := graphql.Do(graphql.Params{
		Schema:        data.Schema,
		RequestString: testutil.IntrospectionQuery,
	})
	if len(result.Errors) > 0 {
		log.Fatalf("❌ Failed introspection: %v", result.Errors)
	}

	jsonBytes, err := json.Marshal(result.Data)
	if err != nil {
		log.Fatalf("❌ Marshal error: %v", err)
	}

	var introspection map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &introspection); err != nil {
		log.Fatalf("❌ Unmarshal error: %v", err)
	}

	sdl := convertIntrospectionToSDL(introspection)

	output := "./server/data/schema.graphql"
	if err := os.WriteFile(output, []byte(sdl), 0644); err != nil {
		log.Fatalf("❌ Write error: %v", err)
	}

	fmt.Println("✅ Schema written to", output)
}
