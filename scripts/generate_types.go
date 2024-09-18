package scripts

import (
	"fmt"
	"strings"
)

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
		return "[]interface{}"
	case "object":
		return "map[string]interface{}"
	default:
		return "interface{}"
	}
}
