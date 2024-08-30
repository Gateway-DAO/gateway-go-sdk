package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"unicode"
)

type SwaggerSchema struct {
	Definitions map[string]struct {
		Type          string                 `json:"type"`
		Properties    map[string]SwaggerType `json:"properties"`
		Enum          []string               `json:"enum"`
		XEnumVarnames []string               `json:"x-enum-varnames"`
	} `json:"definitions"`
	Paths map[string]PathItem `json:"paths"`
}

type PathItem map[string]Operation

type Operation struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	// Add other fields as needed
}

type SwaggerType struct {
	Type string `json:"type"`
	Ref  string `json:"$ref"`
}

func main() {
	data, err := ioutil.ReadFile("api.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var schema SwaggerSchema
	err = json.Unmarshal(data, &schema)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	goTypes := extractGoTypes(schema)

	routeConstants := generateRouteConstants(schema.Paths)

	output := "package main\n\n" + strings.Join(goTypes, "\n\n") + "\n\n" + routeConstants
	err = ioutil.WriteFile("go_types.go", []byte(output), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Go types have been generated and saved to go_types.go")
}

func extractGoTypes(schema SwaggerSchema) []string {
	var goTypes []string

	for defName, def := range schema.Definitions {
		if len(def.Enum) > 0 {
			goTypes = append(goTypes, generateEnum(defName, def.Enum, def.XEnumVarnames))
		} else if def.Type == "object" {
			goTypes = append(goTypes, generateStruct(defName, def.Properties))
		}
	}

	return goTypes
}

func generateEnum(name string, enumValues []string, varNames []string) string {
	typeName := toPascalCase(name)
	var lines []string
	lines = append(lines, fmt.Sprintf("type %s string", typeName))
	lines = append(lines, "")
	lines = append(lines, "const (")

	for i, value := range enumValues {
		var enumName string
		if i < len(varNames) {
			enumName = varNames[i]
		} else {
			enumName = fmt.Sprintf("%s%s", typeName, toPascalCase(value))
		}
		lines = append(lines, fmt.Sprintf("\t%s %s = \"%s\"", enumName, typeName, value))
	}

	lines = append(lines, ")")
	return strings.Join(lines, "\n")
}

func generateStruct(name string, properties map[string]SwaggerType) string {
	var fields []string
	fields = append(fields, fmt.Sprintf("type %s struct {", toPascalCase(name)))

	for propName, propSchema := range properties {
		goType := getGoType(propSchema)
		fields = append(fields, fmt.Sprintf("\t%s %s `json:\"%s\"`", toPascalCase(propName), goType, propName))
	}

	fields = append(fields, "}")
	return strings.Join(fields, "\n")
}

func getGoType(schema SwaggerType) string {
	if schema.Ref != "" {
		parts := strings.Split(schema.Ref, "/")
		return toPascalCase(parts[len(parts)-1])
	}

	switch schema.Type {
	case "integer":
		return "int"
	case "number":
		return "float64"
	case "string":
		return "string"
	case "boolean":
		return "bool"
	case "array":
		return "[]interface{}" // This is a simplification; you might want to handle array item types
	case "object":
		return "map[string]interface{}"
	default:
		return "interface{}"
	}
}

func generateRouteConstants(paths map[string]PathItem) string {
	var constants []string
	constants = append(constants, "const (")

	for path, pathItem := range paths {
		for method, operation := range pathItem {
			constantName := generateRouteConstantName(method, path, operation.Summary)
			parameterizedPath := parameterizePath(path)
			constants = append(constants, fmt.Sprintf("\t%s = \"%s\"", constantName, parameterizedPath))
		}
	}

	constants = append(constants, ")")
	return strings.Join(constants, "\n")
}

func generateRouteConstantName(method, path, summary string) string {
	name := strings.TrimSpace(summary)
	if name == "" {
		name = path
	}
	return toPascalCase(name)
}

func parameterizePath(path string) string {
	re := regexp.MustCompile(`\{([^}]+)\}`)
	return re.ReplaceAllString(path, "%s")
}

func toPascalCase(s string) string {
	var sb strings.Builder
	uppercaseNext := true

	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			uppercaseNext = true
		} else {
			if uppercaseNext {
				sb.WriteRune(unicode.ToUpper(r))
			} else {
				sb.WriteRune(r)
			}
			uppercaseNext = false
		}
	}

	return sb.String()
}
