package scripts

import (
	"fmt"
	"strings"
)

func extractGoTypes(schema SwaggerSchema) []string {
	var goTypes []string

	for defName, def := range schema.Components.Schemas {
		if len(def.Enum) > 0 {
			goTypes = append(goTypes, generateEnum(defName, def.Enum, def.XEnumVarnames))
		} else if def.Type == "object" {
			generatedObject := generateStruct(defName, def.Properties, def.Required)
			if defName == "helper.PaginatedResponse" {
				generatedObject = addGenericsToStruct(generatedObject)
			}
			goTypes = append(goTypes, generatedObject)
		}
	}

	return goTypes
}

func addGenericsToStruct(structDef string) string {
	genericStruct := strings.ReplaceAll(structDef, "interface{}", "T")

	genericStruct = strings.Replace(genericStruct, "struct {", "[T any] struct {", 1)

	return genericStruct
}

func generateEnum(name string, enumValues []string, varNames []string) string {
	typeName := ToPascalCase(name)
	var lines []string
	lines = append(lines, fmt.Sprintf("type %s string", typeName))
	lines = append(lines, "")
	lines = append(lines, "const (")

	for i, value := range enumValues {
		var enumName string
		if i < len(varNames) {
			enumName = varNames[i]
		} else {
			enumName = fmt.Sprintf("%s%s", typeName, ToPascalCase(value))
		}
		lines = append(lines, fmt.Sprintf("\t%s %s = \"%s\"", enumName, typeName, value))
	}

	lines = append(lines, ")")
	return strings.Join(lines, "\n")
}

func generateStruct(name string, properties map[string]SwaggerType, required []string) string {
	var fields []string
	fields = append(fields, fmt.Sprintf("type %s struct {", ToPascalCase(name)))

	requiredSet := make(map[string]struct{})
	for _, field := range required {
		requiredSet[field] = struct{}{}
	}

	for propName, propSchema := range properties {
		goType := getGoType(propSchema)
		if _, isRequired := requiredSet[propName]; !isRequired {
			goType = "*" + goType // Use pointer for optional fields
		}
		fields = append(fields, fmt.Sprintf("\t%s %s `json:\"%s\"`", ToPascalCase(propName), goType, propName))
	}

	fields = append(fields, "}")
	return strings.Join(fields, "\n")
}

func getGoType(schema SwaggerType) string {
	if schema.Ref != "" {
		parts := strings.Split(schema.Ref, "/")
		return ToPascalCase(parts[len(parts)-1])
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
		if schema.Items != nil {
			itemType := getGoType(*schema.Items)
			return "[]" + itemType
		}
		return "[]string"
	case "object":
		return "map[string]interface{}"
	default:
		return "interface{}"
	}
}
